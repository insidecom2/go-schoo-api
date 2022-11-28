package auth

import (
	"net/http"

	"go/go-api/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	Avatar string `json:"avatar" binding:"required"`
}



func Register(c *gin.Context) {
	var json RegisterBody
	if err:= c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var isExits bool= checkUserExist(json)
	if !isExits {
		c.JSON(http.StatusOK, gin.H{"status": false,"message" : "User Exist.",})
		  return
	}
	encryptPassword,_ := bcrypt.GenerateFromPassword([]byte(json.Password), 10)
	userCreate := model.User{Username: json.Username, 
		Password: string(encryptPassword), 
		Fullname: json.Fullname,
		Avatar: json.Avatar}

	model.Db.Create(&userCreate)
    c.JSON(http.StatusCreated, gin.H{
      "status": true,
	  "message" : "User Registered.",
	  "User_id" : userCreate.ID,
    })
}

func checkUserExist(json RegisterBody) bool{
	var userExist model.User
	model.Db.Where("username = ?", json.Username).First(&userExist)
	return userExist.ID <= 0
}



