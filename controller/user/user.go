package user

import (
	Repository "go/go-api/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context){
	userProfile := Repository.GetProfile(c)
	c.JSON(http.StatusOK, gin.H{"status" : true, "message": "OK" , "data":userProfile})
}