package handlers

import (
	"gin-grpc/gateway-service/config"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func ProxyRequest(c *gin.Context, targetURL string) {
	url, err := url.Parse(targetURL)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid target URL"})
		return
	}

	proxyURL := httputil.NewSingleHostReverseProxy(url)
	proxyURL.ServeHTTP(c.Writer, c.Request)
}

func UserProxy(c *gin.Context, targetURL string) {
	ProxyRequest(c, targetURL)
}

func PostProxy(c *gin.Context, targetURL string) {
	ProxyRequest(c, targetURL)
}

func RegisterProxyRoutes(r *gin.Engine, cfg *config.Config) {
	r.Any("/users/*proxyPath", func(c *gin.Context) {
		UserProxy(c, cfg.UserServiceURL)
	})
	r.Any("/posts/*proxyPath", func(c *gin.Context) {
		PostProxy(c, cfg.PostServiceURL)
	})
}
