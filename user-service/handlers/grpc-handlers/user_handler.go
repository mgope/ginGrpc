package handlers

import (
	"context"

	pb "gin-grpc/pb"
	service "gin-grpc/user-service/services"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserGRPCServer struct {
	pb.UnimplementedUserServiceServer
	userService *service.UserService
}

func NewUserGRPCServer(userService *service.UserService) *UserGRPCServer {
	return &UserGRPCServer{
		userService: userService,
	}
}

func (s *UserGRPCServer) GetUserByID(
	ctx context.Context,
	req *pb.GetUserRequest,
) (*pb.GetUserResponse, error) {

	user := s.userService.GetUserByID(req.Id)

	if user == nil {
		return nil, status.Error(codes.NotFound, "User not found")
	}

	return &pb.GetUserResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
