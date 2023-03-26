package servicecontract

import (
	"github.com/fajarcandraaa/implement-gRpc-microservice-protobank/grpc/pb"
)

type GrpcContract struct {
	User pb.UserServiceClient
	Book pb.BookServiceClient
}

func NewGrpcService(user pb.UserServiceClient, book pb.BookServiceClient) *GrpcContract {
	return &GrpcContract{
		User: user,
		Book: book,
	}
}
