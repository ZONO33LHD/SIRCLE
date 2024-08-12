package repository

import (
	"context"

	"github.com/ZONO33LHD/sircle/backend/kakeibo-user-service/internal/domain/model"
)

type UserRepository interface {
	GetUser(ctx context.Context, id string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, user *model.User) error
	SetBudget(ctx context.Context, budget *model.Budget) error
	SetGoal(ctx context.Context, goal *model.Goal) error
	UpdateNotificationPreferences(ctx context.Context, userID string, prefs *model.NotificationPreferences) error
}
