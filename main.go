package main

import (
	adapter "btc-network-monitor/internal/adapter/api/resource"
	"btc-network-monitor/internal/adapter/api/rpc"
	"btc-network-monitor/internal/cronjobs"
	"btc-network-monitor/internal/database"
	"btc-network-monitor/internal/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file")
	}

	database.ConnectDB()
	rpc.NewRPCConfig()

	stop := make(chan struct{})

	// Initialize the cron job scheduler
	go func() {
		cronjobs.InitCronJob(stop)
	}()

	// Start the HTTP server in a separate goroutine
	go func() {
		router := gin.Default()
		handler := adapter.NewHTTPHandler()
		handler.Routes(router)
		port := os.Getenv("PORT")
		if port == "" {
			port = "8088"
		}
		logger.Info(fmt.Sprintf(" Starting server on port %v", port))
		err := router.Run(":" + port)
		if err != nil {
			logger.Error("Failed to start server: " + err.Error())
			stop <- struct{}{}
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig

	logger.Info("Stopping server...")
	// Signal to stop the cron job scheduler
	stop <- struct{}{}
	// Log server stopped
	logger.Info("Server stopped")

}
