package grpcclient

import (
	"log"

	pb "gin-grpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserGrpcClient struct {
	Client pb.UserServiceClient
}

func NewUserGrpcClient() *UserGrpcClient {

	conn, err := grpc.Dial(
		"localhost:50051",
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)

	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewUserServiceClient(conn)

	return &UserGrpcClient{
		Client: client,
	}
}
