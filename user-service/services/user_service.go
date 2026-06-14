package service

import (
	"gin-grpc/user-service/models"
	"gin-grpc/user-service/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetUserByID(id int64) *models.User {
	user := s.repo.GetUserByID(id)
	// if user == nil {
	// 	return nil, fmt.Errorf("user not found")
	// }
	return user
}

func (s *UserService) CreateUser(user models.User) models.User {
	return s.repo.CreateUser(user)
}
