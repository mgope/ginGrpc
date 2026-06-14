package repository

import "gin-grpc/post-service/models"

type PostRepository struct {
	posts []models.Post
}
