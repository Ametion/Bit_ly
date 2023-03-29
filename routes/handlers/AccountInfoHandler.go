package handlers

import (
	"bit-ly/database"
	"bit-ly/models"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AccountInfoHandler(con *gin.Context) {
	var usr database.User

	value, isKeyExist := con.Get("userID")

	if !isKeyExist {
		con.JSON(401, models.ResponseModel{Code: 401, Message: "Can not find user id key, login again"})
	}

	usrID := value.(uint)

	userResult := database.Database.Where("id = ?", usrID).First(&usr)

	if userResult.Error != nil && errors.Is(userResult.Error, gorm.ErrRecordNotFound) {
		con.JSON(404, models.ResponseModel{Code: 400, Message: "Can not find user with presented login"})
		return
	}

	var links []database.ShortLink
	linksResult := database.Database.Where("user_id = ?", usr.ID).Find(&links)

	if linksResult.Error != nil && !errors.Is(linksResult.Error, gorm.ErrRecordNotFound) {
		con.JSON(400, models.ResponseModel{Code: 400, Message: "Error while getting user links"})
		return
	}

	var responseLinks []models.LinksResponse

	for _, link := range links {
		responseLinks = append(responseLinks, models.LinksResponse{OriginalLink: link.OriginalLink,
			ShortedLink: link.ShortedLink})
	}

	con.JSON(200, models.AccountInfoResponse{Links: responseLinks})
}
