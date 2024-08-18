package persistence

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/ZONO33LHD/sircle/backend/kakeibo-user-service/internal/domain/model"
	"github.com/ZONO33LHD/sircle/backend/kakeibo-user-service/internal/repository"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(dataSourceName string) repository.UserRepository {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) GetUser(ctx context.Context, id string) (*model.User, error) {
	query := `SELECT id, name, email, budgets, goals, notification_preferences FROM users WHERE id = $1`
	var user model.User
	var budgetsJSON, goalsJSON, prefsJSON []byte

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID, &user.Name, &user.Email,
		&budgetsJSON, &goalsJSON, &prefsJSON,
	)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(prefsJSON, &user.NotificationPreferences)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepositoryImpl) CreateUser(ctx context.Context, user *model.User) error {
	query := `
        INSERT INTO users (name, email, password_hash, notification_preferences)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `
	prefsJSON, err := json.Marshal(user.NotificationPreferences)
	if err != nil {
		return fmt.Errorf("failed to marshal notification preferences: %w", err)
	}

	err = r.db.QueryRowContext(ctx, query,
		user.Name,
		user.Email,
		user.PasswordHash,
		prefsJSON,
	).Scan(&user.ID)

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (r *UserRepositoryImpl) UpdateUser(ctx context.Context, user *model.User) error {
	query := `UPDATE users SET name = $2, email = $3, notification_preferences = $4
                WHERE id = $1`

	prefsJSON, err := json.Marshal(user.NotificationPreferences)
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, query,
		user.ID, user.Name, user.Email, prefsJSON,
	)
	return err
}

func (r *UserRepositoryImpl) SetBudget(ctx context.Context, budget *model.Budget) error {
	query := `UPDATE users SET budgets = budgets || $2::jsonb WHERE id = $1`
	budgetJSON, err := json.Marshal(budget)
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(ctx, query, budget.UserID, budgetJSON)
	return err
}

func (r *UserRepositoryImpl) SetGoal(ctx context.Context, goal *model.Goal) error {
	query := `UPDATE users SET goals = goals || $2::jsonb WHERE id = $1`
	goalJSON, err := json.Marshal(goal)
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(ctx, query, goal.UserID, goalJSON)
	return err
}

func (r *UserRepositoryImpl) UpdateNotificationPreferences(ctx context.Context, userID string, prefs *model.NotificationPreferences) error {
	query := `UPDATE user SET notification_preferences = $2 WHERE id = $1`
	prefsJSON, err := json.Marshal(prefs)
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(ctx, query, userID, prefsJSON)
	return err
}

func (r *UserRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `SELECT id, name, email, password_hash, notification_preferences FROM users WHERE email = $1`
	var user model.User
	var prefsJSON []byte

	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID, &user.Name, &user.Email, &user.PasswordHash, &prefsJSON,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	err = json.Unmarshal(prefsJSON, &user.NotificationPreferences)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal notification preferences: %w", err)
	}

	return &user, nil
}
