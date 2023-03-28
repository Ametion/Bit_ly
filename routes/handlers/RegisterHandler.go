package handlers

import (
	"bit-ly/database"
	"bit-ly/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(con *gin.Context) {
	var body models.RegisterRequest

	if bodyError := con.ShouldBindJSON(&body); bodyError != nil {
		con.JSON(400, models.ResponseModel{Code: 400, Message: "Error while getting parameters from body"})
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	if err != nil {
		panic("error while hashing")
	}

	usr := database.User{
		FirstName:  body.FirstName,
		SecondName: body.SecondName,
		Login:      body.Login,
		Password:   string(hashedPass),
	}

	database.Database.Create(&usr)
	con.JSON(201, models.ResponseModel{Code: 201, Message: "User created"})
}
