package service

import (
	"context"

	pb "gin-grpc/pb"
)

type PostService struct {
	userClient pb.UserServiceClient
}

func NewPostService(
	client pb.UserServiceClient,
) *PostService {

	return &PostService{
		userClient: client,
	}
}

func (s *PostService) GetPostAuthor(
	userID int64,
) (*pb.GetUserResponse, error) {

	return s.userClient.GetUserByID(
		context.Background(),
		&pb.GetUserRequest{
			Id: userID,
		},
	)
}
