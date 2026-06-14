package routes

import (
	httpHandler "gin-grpc/user-service/handlers/http-handlers"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(
	r *gin.Engine,
	h *httpHandler.UserHandler,
) {
	userGroup := r.Group("/users")
	userGroup.GET("/:id", h.GetUserByID)
	userGroup.POST("/create-user", h.CreateUser)
}
