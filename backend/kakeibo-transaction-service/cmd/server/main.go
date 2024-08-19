package main

import (
	"fmt"
	"log"
	"os"

	transactiongrpc "github.com/ZONO33LHD/sircle/backend/kakeibo-transaction-service/pkg/grpc"
	"github.com/ZONO33LHD/sircle/backend/kakeibo-transaction-service/infrastructure/persistence"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 環境変数から接続情報を取得
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSLMODE")

	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, sslMode)

	transactionRepo := persistence.NewTransactionRepository(dataSourceName)

	transactionService := transactiongrpc.NewTransactionService(transactionRepo)

	server := transactiongrpc.NewServer(transactionService)

	log.Println("Starting gRPC server on :50052")
	if err := server.Run(":50052"); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
