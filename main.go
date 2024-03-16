package main

import (
	adapter "btc-network-monitor/internal/adapter/api/resource"

	"btc-network-monitor/internal/database"
	"btc-network-monitor/internal/logger"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file")
	}

	database.ConnectDB()

	router := gin.Default()
	handler := adapter.NewHTTPHandler()
	handler.Routes(router)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8088"
	}
	logger.Info(fmt.Sprintf(" Starting server on port %v", port))
	router.Run(":" + port)
}
