package persistence

import (
	"context"

	"github.com/ZONO33LHD/sircle/backend/kakeibo-user-service/internal/domain/model"
	"github.com/ZONO33LHD/sircle/backend/kakeibo-user-service/internal/repository"
)

type UserRepositoryImpl struct {
	// ここにデータベース接続などの必要なフィールドを追加
}

func NewUserRepository() repository.UserRepository {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) GetUser(ctx context.Context, id string) (*model.User, error) {
	// TODO: 実装
	return nil, nil
}

func (r *UserRepositoryImpl) CreateUser(ctx context.Context, user *model.User) error {
	// TODO: 実装
	return nil
}

func (r *UserRepositoryImpl) UpdateUser(ctx context.Context, user *model.User) error {
	// TODO: 実装
	return nil
}

func (r *UserRepositoryImpl) SetBudget(ctx context.Context, budget *model.Budget) error {
	// TODO: 実装
	return nil
}

func (r *UserRepositoryImpl) SetGoal(ctx context.Context, goal *model.Goal) error {
	// TODO: 実装
	return nil
}

func (r *UserRepositoryImpl) UpdateNotificationPreferences(ctx context.Context, userID string, prefs *model.NotificationPreferences) error {
	// TODO: 実装
	return nil
}
