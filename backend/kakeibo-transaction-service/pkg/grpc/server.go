package grpc

import (
	"log"
	"net"

	"github.com/ZONO33LHD/sircle/backend/kakeibo-transaction-service/pkg/grpc/pb"
	"google.golang.org/grpc"
)

type Server struct {
	transactionService *TransactionService
}

func NewServer(transactionService *TransactionService) *Server {
	return &Server{
		transactionService: transactionService,
	}
}

func (s *Server) Run(port string) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTransactionServiceServer(grpcServer, s.transactionService)

	log.Printf("gRPC server listening on %s", port)
	return grpcServer.Serve(lis)
}
