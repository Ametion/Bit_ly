package handlers

import (
	"bit-ly/internal/database"
	database_models "bit-ly/internal/database/models"
	"bit-ly/internal/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

func CreateLinkHandler(context *gin.Context) {
	var body models.CreateLinkRequest

	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(400, models.Response{Code: 400, Message: "Error while getting parameters from body"})
		return
	}

	shortLink, err := findUniqueShortLink()

	if err != nil {
		context.JSON(400, models.Response{Code: 400, Message: "Error while creating your short link"})
		return
	}

	link := database_models.ShortLink{ShortLink: shortLink, OriginalLink: body.OriginalLink}

	database.GetConnection().Create(&link)
	context.JSON(201, models.Response{Code: 201, Message: fmt.Sprintf("Your link is ready - localhost:5000/%s", link.ShortLink)})
}

func findUniqueShortLink() (string, error) {
	var shortLink string
	var oldLink database_models.ShortLink

	for {
		shortLink = generateRandomString(10)
		result := database.GetConnection().Where("short_link = ?", shortLink).First(&oldLink)

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
