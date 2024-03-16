package database

import (
	"fmt"
	"os"

	"github.com/subosito/gotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBInstance *gorm.DB

func ConnectDB() *gorm.DB {
	_ = gotenv.Load()

	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?%v",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DATABASE"),
		os.Getenv("MYSQL_DEFAULT_CONFIG"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}

	DBInstance = db

	return db
}
