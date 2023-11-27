package main

import (
	"bit-ly/internal/database"
	"bit-ly/internal/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	database.ConnectDatabase()

	r := gin.Default()

	r.POST("/link", handlers.CreateLinkHandler)
	r.GET("/:link", handlers.RedirectLinkHandler)

	runErr := r.Run(":5000")
	if runErr != nil {
		log.Fatalf("Errror while running: %s", runErr.Error())
	}
}
