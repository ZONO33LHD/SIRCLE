package persistence

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/ZONO33LHD/sircle/backend/kakeibo-transaction-service/domain/model"
	"github.com/ZONO33LHD/sircle/backend/kakeibo-transaction-service/repository"
)

type TransactionRepositoryImpl struct {
	db *sql.DB
}

func NewTransactionRepository(dataSourceName string) repository.TransactionRepository {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return &TransactionRepositoryImpl{db: db}
}

func (r *TransactionRepositoryImpl) GetTransaction(ctx context.Context, id string) (*model.Transaction, error) {
	// TODO: 実装を追加
	return nil, nil
}

func (r *TransactionRepositoryImpl) GetTransactionsByCategoryId(ctx context.Context, categoryId string) ([]*model.Transaction, error) {
	// TODO: 実装を追加
	return nil, nil
}

func (r *TransactionRepositoryImpl) GetTransactionsByDateRange(ctx context.Context, startDate, endDate time.Time) ([]*model.Transaction, error) {
	// TODO: 実装を追加
	return nil, nil
}

func (r *TransactionRepositoryImpl) GetTransactionsByType(ctx context.Context, transactionType string) ([]*model.Transaction, error) {
	// TODO: 実装を追加
	return nil, nil
}

func (r *TransactionRepositoryImpl) GetTransactionsByUserId(ctx context.Context, userId string) ([]*model.Transaction, error) {
	// TODO: 実装を追加
	return nil, nil
}

func (r *TransactionRepositoryImpl) UpdateTransaction(ctx context.Context, transaction *model.Transaction) (*model.Transaction, error) {

	return nil, nil
}

func (r *TransactionRepositoryImpl) CreateTransaction(ctx context.Context, transaction *model.Transaction) (*model.Transaction, error) {
	query := `
		INSERT INTO transactions (user_id, amount, type, category_id, date, description, is_recurring, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id
	`
	now := time.Now()
	var id int
	err := r.db.QueryRowContext(ctx, query,
		transaction.UserId,
		transaction.Amount,
		transaction.Type,
		transaction.CategoryId,
		transaction.Date,
		transaction.Description,
		transaction.IsRecurring,
		now,
		now,
	).Scan(&id)

	if err != nil {
		return nil, fmt.Errorf("取引の作成に失敗しました: %v", err)
	}

	// カテゴリー情報を取得
	var categoryName string
	err = r.db.QueryRowContext(ctx, "SELECT name FROM categories WHERE id = $1", transaction.CategoryId).Scan(&categoryName)
	if err != nil {
		return nil, fmt.Errorf("カテゴリー情報の取得に失敗しました: %v", err)
	}

	transaction.ID = strconv.Itoa(id)
	transaction.CategoryName = categoryName
	transaction.CreatedAt = now
	transaction.UpdatedAt = now
	return transaction, nil
}

func (r *TransactionRepositoryImpl) DeleteTransaction(ctx context.Context, id string) error {
	query := `DELETE FROM transactions WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

