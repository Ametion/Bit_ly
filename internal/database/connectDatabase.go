package database

import (
	"bit-ly/internal/database/models"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var database *gorm.DB

func ConnectDatabase() {
	errEnv := godotenv.Load()

	if errEnv != nil {
		fmt.Println("Error loading .env file")
	}

	dns := os.Getenv("databaseConnection")

	var err error
	database, err = gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	migrateErr := database.AutoMigrate(&database_models.ShortLink{})

	if migrateErr != nil {
		panic(migrateErr.Error())
	}
}

func GetConnection() *gorm.DB {
	return database
}
