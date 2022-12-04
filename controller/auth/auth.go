package auth

import (
	"go/go-api/model"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type RegisterBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	Avatar string `json:"avatar" binding:"required"`
}

type LoginBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
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

func Login(c *gin.Context) {
	var json LoginBody
	if err:= c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var isExits model.User= checkLoginExist(json)
	if isExits.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": false,"message" : "User Does No Exist."})
		return
	}
	isPasswordMatch := IsPasswordMatch(isExits.Password, json.Password)
	if !isPasswordMatch {
		c.JSON(http.StatusBadRequest, gin.H{"status": false,"message" : "Password wrong."})
		return
	}
	token := getJWTToken(isExits)
	c.JSON(http.StatusAccepted, gin.H{"status": true,"message" : "ok", "token": token})

}

var hmacSampleSecret []byte
func getJWTToken(user model.User) string {
	secret := os.Getenv("JWT_SECRET")
	hmacSampleSecret = []byte(secret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId" : user.ID,
		"username" : user.Username,
		"exp" : time.Now().Add(time.Minute*120).Unix(),
		})
	tokenString, _ := token.SignedString(hmacSampleSecret)
	
	return tokenString
}

func checkUserExist(json RegisterBody) bool{
	var userExist model.User
	model.Db.Where("username = ?", json.Username).First(&userExist)
	return userExist.ID <= 0
}

func checkLoginExist(json LoginBody) model.User{
	var userExist model.User
	model.Db.Where("username = ?", json.Username).First(&userExist)
	return userExist
}

func IsPasswordMatch(passwordHash string, password string) bool {
	result := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return result == nil
}


