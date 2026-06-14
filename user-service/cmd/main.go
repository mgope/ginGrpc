package main

import (
	"log"
	"net"

	pb "gin-grpc/pb"
	grpcHandler "gin-grpc/user-service/handlers/grpc-handlers"
	httpHandler "gin-grpc/user-service/handlers/http-handlers"
	"gin-grpc/user-service/repository"
	"gin-grpc/user-service/routes"
	service "gin-grpc/user-service/services"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func main() {
	// Dependencies
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)

	userHandler := httpHandler.NewUserHandler(userService)
	userGRPCServer := grpcHandler.NewUserGRPCServer(userService)

	// gRPC Server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userGRPCServer)

	// HTTP Server
	router := gin.Default()
	routes.SetupUserRoutes(router, userHandler)

	var g errgroup.Group

	g.Go(func() error {
		log.Println("gRPC server running on :50051")
		return grpcServer.Serve(lis)
	})

	g.Go(func() error {
		log.Println("HTTP server running on :8080")
		return router.Run(":8080")
	})

	if err := g.Wait(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
