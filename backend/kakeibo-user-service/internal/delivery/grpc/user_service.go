package grpc

import (
	"context"

	pb "github.com/ZONO33LHD/sircle/backend/kakeibo-user-service/internal/delivery/grpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	// TODO: 実際のユーザー取得ロジックを実装する
	return nil, status.Errorf(codes.Unimplemented, "GetUser not implemented yet")
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	// TODO: 実際のユーザー作成ロジックを実装する
	return nil, status.Errorf(codes.Unimplemented, "CreateUser not implemented yet")
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.User, error) {
	// TODO: 実際のユーザー更新ロジックを実装する
	return nil, status.Errorf(codes.Unimplemented, "UpdateUser not implemented yet")
}

func (s *UserService) SetBudget(ctx context.Context, req *pb.SetBudgetRequest) (*pb.Budget, error) {
	// TODO: 実際の予算設定ロジックを実装する
	return nil, status.Errorf(codes.Unimplemented, "SetBudget not implemented yet")
}

func (s *UserService) SetGoal(ctx context.Context, req *pb.SetGoalRequest) (*pb.Goal, error) {
	// TODO: 実際の目標設定ロジックを実装する
	return nil, status.Errorf(codes.Unimplemented, "SetGoal not implemented yet")
}

func (s *UserService) UpdateNotificationPreferences(ctx context.Context, req *pb.UpdateNotificationPreferencesRequest) (*pb.NotificationPreferences, error) {
	// TODO: 実際の通知設定更新ロジックを実装する
	return nil, status.Errorf(codes.Unimplemented, "UpdateNotificationPreferences not implemented yet")
}
