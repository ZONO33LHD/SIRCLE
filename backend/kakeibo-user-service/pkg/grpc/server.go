package grpc

import (
	"log"
	"net"

	pb "github.com/ZONO33LHD/sircle/backend/kakeibo-user-service/pkg/grpc/pb"
	"google.golang.org/grpc"
)

type Server struct {
	userService *UserService
}

func NewServer(userService *UserService) *Server {
	return &Server{
		userService: userService,
	}
}

func (s *Server) Run(port string) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, s.userService)

	log.Printf("gRPC server listening on %s", port)
	return grpcServer.Serve(lis)
}
