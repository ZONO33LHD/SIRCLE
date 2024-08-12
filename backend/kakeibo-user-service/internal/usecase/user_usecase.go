package usecase

import (
	"context"

	"github.com/ZONO33LHD/sircle/backend/kakeibo-user-service/internal/domain/model"
	"github.com/ZONO33LHD/sircle/backend/kakeibo-user-service/internal/repository"
)

type UserUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *UserUsecase {
	return &UserUsecase{userRepo: userRepo}
}

func (u *UserUsecase) GetUser(ctx context.Context, id string) (*model.User, error) {
	return u.userRepo.GetUser(ctx, id)
}

func (u *UserUsecase) CreateUser(ctx context.Context, user *model.User) error {
	return u.userRepo.CreateUser(ctx, user)
}

func (u *UserUsecase) UpdateUser(ctx context.Context, user *model.User) error {
	return u.userRepo.UpdateUser(ctx, user)
}

func (u *UserUsecase) SetBudget(ctx context.Context, budget *model.Budget) error {
	return u.userRepo.SetBudget(ctx, budget)
}

func (u *UserUsecase) SetGoal(ctx context.Context, goal *model.Goal) error {
	return u.userRepo.SetGoal(ctx, goal)
}

func (u *UserUsecase) UpdateNotificationPreferences(ctx context.Context, userID string, prefs *model.NotificationPreferences) error {
	return u.userRepo.UpdateNotificationPreferences(ctx, userID, prefs)
}
