package handlers

import (
	"bit-ly/database"
	"bit-ly/models"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"os"
	"time"
)

func LoginHandler(con *gin.Context) {
	var body models.LoginRequest

	if bodyError := con.ShouldBindJSON(&body); bodyError != nil {
		con.JSON(400, models.ResponseModel{Code: 400, Message: "Error while getting parameters from body"})
		return
	}

	var usr database.User

	result := database.Database.Where("login = ?", body.Login).First(&usr)

	if result != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		con.JSON(404, models.ResponseModel{Code: 404, Message: "Can not find user with presented login"})
		return
	}

	hashError := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(body.Password))

	if hashError != nil {
		con.JSON(400, models.ResponseModel{Code: 400, Message: "Error while check password"})
		return
	}

	token, tokenError := GenerateToken(&usr)

	if tokenError != nil {
		con.JSON(400, models.ResponseModel{Code: 400, Message: "Error while logged in"})
		return
	}

	con.JSON(200, models.LoginResponse{Token: token})
}

func GenerateToken(usr *database.User) (string, error) {
	errEnv := godotenv.Load()

	if errEnv != nil {
		fmt.Println("Error loading .env file")
	}

	expirationTime := time.Now().Add(time.Minute)

	claims := &Claims{
		UserID: usr.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("jwtSecret")))
}

type Claims struct {
	UserID uint
	jwt.StandardClaims
}
