package repository

import "gin-grpc/user-service/models"

type UserRepository struct {
	users []models.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: []models.User{},
	}
}

func (r *UserRepository) GetUserByID(id int64) *models.User {
	for _, user := range r.users {
		if user.ID == id {
			return &user
		}
	}
	return nil
}

func (r *UserRepository) CreateUser(user models.User) models.User {
	user.ID = int64(len(r.users) + 1)
	r.users = append(r.users, user)
	return user
}
