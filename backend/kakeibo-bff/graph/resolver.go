package graph

import (
	pb "github.com/ZONO33LHD/sircle/backend/kakeibo-user-service/pkg/grpc/pb"
	"google.golang.org/grpc"
)

type Resolver struct{
	UserServiceClient pb.UserServiceClient
}

func NewResolver(conn *grpc.ClientConn) *Resolver {
	return &Resolver{
		UserServiceClient: pb.NewUserServiceClient(conn),
	}
}