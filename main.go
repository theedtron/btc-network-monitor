package main

import (
	adapter "btc-network-monitor/internal/adapter/api/resource"
	"btc-network-monitor/internal/adapter/api/rpc"
	"btc-network-monitor/internal/cronjobs"
	"time"

	"btc-network-monitor/internal/database"
	"btc-network-monitor/internal/logger"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/robfig/cron"

)

func main() {

	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file")
	}

	database.ConnectDB()
	rpc.NewRPCConfig()

	router := gin.Default()
	handler := adapter.NewHTTPHandler()
	handler.Routes(router)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8088"
	}
	logger.Info(fmt.Sprintf(" Starting server on port %v", port))
	router.Run(":" + port)

	//Add cron job
	c := cron.New()
	// Define the Cron job schedule
    c.AddFunc("30 * * * *", func() {
        cronjobs.TxNotify()
    })
	// Start the Cron job scheduler
    c.Start()
	// Wait for the Cron job to run
    time.Sleep(5 * time.Minute)

    // Stop the Cron job scheduler
    c.Stop()
}
