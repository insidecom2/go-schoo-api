package repository

import (
	"go/go-api/model"

	"github.com/gin-gonic/gin"
)
type responseUserProFile struct {
	ID uint
	Username string 
	Fullname string 
	Avatar string 
}

func GetProfile (c *gin.Context) responseUserProFile{
	var user model.User
	userId := c.MustGet("userId").(float64)
	model.Db.First(&user, userId)

	userResponse := responseUserProFile {
		ID: user.ID,
		Username: user.Username, 
		Fullname: user.Fullname,
		Avatar: user.Avatar,
	}

	return 	userResponse

}