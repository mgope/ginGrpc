package main

import (
	grpcclient "gin-grpc/post-service/grpc"
	handler "gin-grpc/post-service/handlers/http-handlers"
	service "gin-grpc/post-service/services"

	"github.com/gin-gonic/gin"

	routes "gin-grpc/post-service/routes"
)

func main() {

	userGrpc :=
		grpcclient.NewUserGrpcClient()

	postService :=
		service.NewPostService(
			userGrpc.Client,
		)

	postHandler :=
		handler.NewPostHandler(
			postService,
		)

	r := gin.Default()

	routes.SetupRoutes(
		r,
		postHandler,
	)

	r.Run(":8081")
}
