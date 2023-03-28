package handlers

import (
	"bit-ly/database"
	"bit-ly/models"
	"github.com/gin-gonic/gin"
)

func RedirectLinkHandler(con *gin.Context) {
	shortLink := con.Param("link")

	var link database.ShortLink

	result := database.Database.Where("short_link = ?", shortLink).First(&link)

	if result.Error != nil {
		con.JSON(404, models.ResponseModel{Code: 404, Message: "Can not find any link with this shortcut"})
		return
	}

	con.Redirect(302, link.OriginalLink)
}
