package handlers

import (
	"bit-ly/database"
	"bit-ly/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

func CreateLinkHandler(con *gin.Context) {
	var body models.CreateLinkRequest

	if err := con.ShouldBindJSON(&body); err != nil {
		con.JSON(400, models.ResponseModel{Code: 400, Message: "Error while getting parameters from body"})
		return
	}

	shortLink, err := findUniqueShortLink()

	if err != nil {
		con.JSON(400, models.ResponseModel{Code: 400, Message: "Error while creating your short link"})
		return
	}

	link := database.ShortLink{ShortedLink: shortLink, OriginalLink: body.OriginalLink}

	database.Database.Create(&link)
	con.JSON(201, models.ResponseModel{Code: 201, Message: fmt.Sprintf("Your link is ready - localhost:5000/%s", link.ShortedLink)})
}

func findUniqueShortLink() (string, error) {
	var shortLink string
	var oldLink database.ShortLink

	for {
		shortLink = generateRandomString(10)
		result := database.Database.Where("short_link = ?", shortLink).First(&oldLink)

		if result.Error != nil {
			if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return "", errors.New("database error")
			} else {
				break
			}
		}
	}

	return shortLink, nil
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		result[i] = charset[rand.Intn(len(charset))]
	}

	return string(result)
}
