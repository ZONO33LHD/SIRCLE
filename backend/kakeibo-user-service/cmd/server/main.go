package main

import (
	"log"

	usergrpc "github.com/ZONO33LHD/sircle/backend/kakeibo-user-service/internal/delivery/grpc"
)

func main() {
	userService := usergrpc.NewUserService()
	server := usergrpc.NewServer(userService)

	log.Println("Starting gRPC server on :50051")
	if err := server.Run(":50051"); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
