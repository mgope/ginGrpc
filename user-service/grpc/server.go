package grpcserver

import (
	"context"
	"fmt"
	"gin-grpc/pb"
	service "gin-grpc/user-service/services"
)

type UserGrpcServer struct {
	pb.UnimplementedUserServiceServer
	userService *service.UserService
}

func NewUserGrpcServer(userService *service.UserService) *UserGrpcServer {
	return &UserGrpcServer{
		userService: userService,
	}
}

func (s *UserGrpcServer) GetUserByID(
	ctx context.Context,
	req *pb.GetUserRequest,
) (*pb.GetUserResponse, error) {
	user := s.userService.GetUserByID(req.Id)
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	return &pb.GetUserResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
