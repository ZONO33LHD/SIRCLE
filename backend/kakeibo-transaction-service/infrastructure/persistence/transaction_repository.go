package persistence

import (

	"database/sql"
	"log"

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