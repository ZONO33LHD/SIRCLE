package repository

import (
	"context"
	"time"

	"github.com/ZONO33LHD/sircle/backend/kakeibo-transaction-service/domain/model"
)

type TransactionRepository interface {
	GetTransaction(ctx context.Context, id string) (*model.Transaction, error)
	CreateTransaction(ctx context.Context, transaction *model.Transaction) (*model.Transaction, error)
	UpdateTransaction(ctx context.Context, transaction *model.Transaction) (*model.Transaction, error)
	DeleteTransaction(ctx context.Context, id string) error
	GetTransactionsByUserId(ctx context.Context, userId string) ([]*model.Transaction, error)
	GetTransactionsByCategoryId(ctx context.Context, categoryId string) ([]*model.Transaction, error)
	GetTransactionsByDateRange(ctx context.Context, startDate time.Time, endDate time.Time) ([]*model.Transaction, error)
	GetTransactionsByType(ctx context.Context, transactionType string) ([]*model.Transaction, error)
}
