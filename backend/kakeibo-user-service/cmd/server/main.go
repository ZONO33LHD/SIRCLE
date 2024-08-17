package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ZONO33LHD/sircle/backend/kakeibo-user-service/internal/infrastructure/persistence"
	usergrpc "github.com/ZONO33LHD/sircle/backend/kakeibo-user-service/pkg/grpc"
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
	userRepo := persistence.NewUserRepository(dataSourceName)
	userService := usergrpc.NewUserService(userRepo)
	server := usergrpc.NewServer(userService)

	log.Println("Starting gRPC server on :50051")
	if err := server.Run(":50051"); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
