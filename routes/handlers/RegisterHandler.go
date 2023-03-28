package handlers

import (
	"bit-ly/database"
	"bit-ly/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(con *gin.Context) {
	var registerBody models.RegisterRequest

	if bodyError := con.ShouldBindJSON(&registerBody); bodyError != nil {
		con.JSON(400, models.ResponseModel{Code: 400, Message: "Error while getting parameters from body"})
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(registerBody.Password), bcrypt.DefaultCost)

	if err != nil {
		panic("error while hashing")
	}

	usr := database.User{
		FirstName:  registerBody.FirstName,
		SecondName: registerBody.SecondName,
		Login:      registerBody.Login,
		Password:   string(hashedPass),
	}

	database.Database.Create(&usr)
	con.JSON(201, models.ResponseModel{Code: 201, Message: "User created"})
}
