package grpc

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/ZONO33LHD/sircle/backend/kakeibo-user-service/internal/domain/model"
	"github.com/ZONO33LHD/sircle/backend/kakeibo-user-service/internal/repository"
	pb "github.com/ZONO33LHD/sircle/backend/kakeibo-user-service/pkg/grpc/pb"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	// TODO: 実際のユーザー取得ロジックを実装する
	return nil, status.Errorf(codes.Unimplemented, "GetUser not implemented yet")
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "パスワードのハッシュ化に失敗しました: %v", err)
	}

	user := &model.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		NotificationPreferences: &model.NotificationPreferences{
			BudgetNotifications: false,
			GoalNotifications:   false,
		},
	}

	err = s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "ユーザーの作成に失敗しました: %v", err)
	}

	return &pb.User{
		Id:    strconv.Itoa(user.ID),
		Name:  user.Name,
		Email: user.Email,
		NotificationPreferences: &pb.NotificationPreferences{
			BudgetNotifications: user.NotificationPreferences.BudgetNotifications,
			GoalNotifications:   user.NotificationPreferences.GoalNotifications,
		},
	}, nil
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

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Printf("Login called with email: %s", req.Email)
	user, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	log.Printf("GetUserByEmail result: %+v", user)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "ユーザーが見つかりません")
	}

	// パスワードの比較
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "無効なパスワードです")
	}

	// JWTトークンの生成
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 環境変数から秘密鍵を取得
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "トークンの生成に失敗しました")
	}

	return &pb.LoginResponse{
		User: &pb.User{
			Id:    strconv.Itoa(user.ID),
			Name:  user.Name,
			Email: user.Email,
		},
		Token: tokenString,
	}, nil
}
