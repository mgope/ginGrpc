package handler

import (
	"net/http"
	"strconv"

	service "gin-grpc/post-service/services"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	service *service.PostService
}

func NewPostHandler(
	service *service.PostService,
) *PostHandler {

	return &PostHandler{
		service: service,
	}
}

func (h *PostHandler) GetPost(
	c *gin.Context,
) {

	id, _ := strconv.ParseInt(
		c.Param("id"),
		10,
		64,
	)

	user, err :=
		h.service.GetPostAuthor(id)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"post": "grpc demo",
			"user": user,
		},
	)
}
