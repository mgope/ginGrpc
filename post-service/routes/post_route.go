package routes

import (
	httpHandler "gin-grpc/post-service/handlers/http-handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	r *gin.Engine,
	h *httpHandler.PostHandler,
) {

	postGroup := r.Group("/posts")

	postGroup.GET("/:id", h.GetPost)
}
