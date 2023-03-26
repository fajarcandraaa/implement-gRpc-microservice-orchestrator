package client

import (
	"log"
	"os"

	"github.com/fajarcandraaa/implement-gRpc-microservice-protobank/grpc/pb"
	"google.golang.org/grpc"
)

func ServiceUser() pb.UserServiceClient {
	port := os.Getenv("GRPC_CLIENT_USER")
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return pb.NewUserServiceClient(conn)
}

func ServiceBook() pb.BookServiceClient {
	port := os.Getenv("GRPC_CLIENT_BOOK")
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return pb.NewBookServiceClient(conn)
}
