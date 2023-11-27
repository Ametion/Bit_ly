package handlers

import (
	"bit-ly/internal/database"
	database_models "bit-ly/internal/database/models"
	"bit-ly/internal/models"
	"github.com/gin-gonic/gin"
)

func RedirectLinkHandler(context *gin.Context) {
	shortLink := context.Param("link")

	var link database_models.ShortLink

	result := database.GetConnection().Where("short_link = ?", shortLink).First(&link)

	if result.Error != nil {
		context.JSON(404, models.Response{Code: 404, Message: "Can not find any link with this shortcut"})
		return
	}

	context.Redirect(302, link.OriginalLink)
}
