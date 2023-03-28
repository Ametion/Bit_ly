package middleware

import (
	"bit-ly/routes/handlers"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		tokenString := bearerToken[1]

		token, err := jwt.ParseWithClaims(tokenString, &handlers.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("jwtSecret")), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Error parsing token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*handlers.Claims); ok && token.Valid {
			c.Set("userId", claims.UserID)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
	}
}
