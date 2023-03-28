package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var Database *gorm.DB

func ConnectDatabase() {
	errEnv := godotenv.Load()

	if errEnv != nil {
		fmt.Println("Error loading .env file")
	}

	dns := os.Getenv("databaseConnection")

	var err error
	Database, err = gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	migrateErr := Database.AutoMigrate(&ShortLink{})

	if migrateErr != nil {
		panic(migrateErr.Error())
	}
}
