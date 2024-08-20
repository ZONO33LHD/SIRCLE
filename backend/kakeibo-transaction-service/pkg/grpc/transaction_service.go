package grpc

import (
	"context"
	"fmt"
	"time"

	"github.com/ZONO33LHD/sircle/backend/kakeibo-transaction-service/domain/model"
	"github.com/ZONO33LHD/sircle/backend/kakeibo-transaction-service/pkg/grpc/pb"
	"github.com/ZONO33LHD/sircle/backend/kakeibo-transaction-service/repository"
)

type TransactionService struct {
	pb.UnimplementedTransactionServiceServer
	repo repository.TransactionRepository
}

func NewTransactionService(repo repository.TransactionRepository) *TransactionService {
	return &TransactionService{
		repo: repo,
	}
}

func (s *TransactionService) CreateTransaction(ctx context.Context, req *pb.CreateTransactionRequest) (*pb.CreateTransactionResponse, error) {
	date, err := time.Parse(time.RFC3339, req.Date)
	if err != nil {
		return nil, fmt.Errorf("無効な日付形式です: %v", err)
	}

	transaction := &model.Transaction{
		UserId:       req.UserId,
		Amount:       float64(req.Amount),
		Type:         req.Type,
		CategoryId:   req.CategoryId,
		CategoryName: req.CategoryName,
		Date:         date,
		Description:  req.Description,
		IsRecurring:  req.IsRecurring,
	}

	createdTransaction, err := s.repo.CreateTransaction(ctx, transaction)
	if err != nil {
		return nil, err
	}

	return &pb.CreateTransactionResponse{
		Transaction: &pb.Transaction{
			Id:           createdTransaction.ID,
			UserId:       createdTransaction.UserId,
			Amount:       float32(createdTransaction.Amount),
			Type:         createdTransaction.Type,
			CategoryId:   createdTransaction.CategoryId,
			CategoryName: createdTransaction.CategoryName,
			Date:         createdTransaction.Date.Format(time.RFC3339),
			Description:  createdTransaction.Description,
			IsRecurring:  createdTransaction.IsRecurring,
		},
	}, nil
}
