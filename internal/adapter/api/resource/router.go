package resource

import (
	"btc-network-monitor/internal/adapter/api/middleware"

	"github.com/gin-gonic/gin"
)

func (s *HTTPHandler) Routes(router *gin.Engine) {
	router.Use(middleware.CORS(), middleware.LoggingMiddleWare())
	api := router.Group("/api/v1")

	api.GET("/healthcheck", s.HealthCheck)
	api.POST("/register", s.Register)
	api.POST("/login", s.Login)
	api.GET("/info", s.GetBlockchainInfo)
	api.GET("/block/:block_hash", s.GetBlockByHash)
	api.GET("/blocks/latest", s.GetLatestBlocks)
	api.GET("/block/height/:param", s.GetBlockByHeight)
	api.GET("/transaction/:tx_id", s.GetTransaction)
	api.GET("/transactions/latest", s.GetLatestTransactions)

	//user routers
	users := api.Group("/users")
	{
		users.Use(middleware.AuthMiddleware())
		users.GET("/", s.GetAllUsers)
		users.PUT("/:id", s.UpdateUser)
		users.GET("/:id", s.FindUser)
	}

	router.NoRoute(func(c *gin.Context) { c.String(404, "Not found") })
}
