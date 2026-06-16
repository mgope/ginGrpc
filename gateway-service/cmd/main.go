package main

import (
	"gin-grpc/gateway-service/config"
	"gin-grpc/gateway-service/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	r := gin.Default()

	handlers.RegisterProxyRoutes(r, cfg)

	r.Run(":" + cfg.Port)
}
