package routes

import (
	"bit-ly/routes/handlers"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes() {
	r := gin.Default()

	r.POST("/link", handlers.CreateLinkHandler)
	r.GET("/:link", handlers.RedirectLinkHandler)

	r.Run(":5000")
}
