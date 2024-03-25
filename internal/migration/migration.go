package main

import (
	"btc-network-monitor/internal/core/domain"
	"btc-network-monitor/internal/database"
)

func Migrate() {
	db := database.ConnectDB()
	db.AutoMigrate(&domain.User{}, &domain.TxSubscribe{})
}

func main() {
	Migrate()
}
