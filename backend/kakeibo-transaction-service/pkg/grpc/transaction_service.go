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

func (s *TransactionService) GetIncomeExpenseSummary(ctx context.Context, req *pb.GetIncomeExpenseSummaryRequest) (*pb.GetIncomeExpenseSummaryResponse, error) {
	startDate, err := time.Parse(time.RFC3339, req.StartDate)
	if err != nil {
		return nil, fmt.Errorf("無効な開始日付形式です: %v", err)
	}

	endDate, err := time.Parse(time.RFC3339, req.EndDate)
	if err != nil {
		return nil, fmt.Errorf("無効な終了日付形式です: %v", err)
	}

	summary, err := s.repo.GetIncomeExpenseSummary(ctx, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("収支サマリーの取得に失敗しました: %v", err)
	}

	pbIncomeItems := make([]*pb.SummaryItem, len(summary.IncomeItems))
	for i, item := range summary.IncomeItems {
		pbIncomeItems[i] = &pb.SummaryItem{
			Title:  item.Title,
			Amount: float32(item.Amount),
		}
	}

	pbExpenseItems := make([]*pb.SummaryItem, len(summary.ExpenseItems))
	for i, item := range summary.ExpenseItems {
		pbExpenseItems[i] = &pb.SummaryItem{
			Title:  item.Title,
			Amount: float32(item.Amount),
		}
	}

	return &pb.GetIncomeExpenseSummaryResponse{
		IncomeItems:  pbIncomeItems,
		ExpenseItems: pbExpenseItems,
		Balance:      float32(summary.Balance),
	}, nil
}
