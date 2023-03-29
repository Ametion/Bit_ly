package routes

import (
	"bit-ly/middleware"
	"bit-ly/routes/handlers"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes() {
	r := gin.Default()

	r.GET("/:link", handlers.RedirectLinkHandler)
	r.POST("/login", handlers.LoginHandler)
	r.POST("/register", handlers.RegisterHandler)

	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.POST("/link", handlers.CreateLinkHandler)
		authorized.GET("/account", handlers.AccountInfoHandler)
	}

	r.Run(":5000")
}
