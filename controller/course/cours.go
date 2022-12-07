package course

import (
	"go/go-api/model"
	Repository "go/go-api/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type validateCourse struct {
	Name string `json:"name" binding:"required"`
}

func CreateCourse(c *gin.Context){
	var json validateCourse
	if err:= c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	courseCreate :=  model.Course{
		Name: json.Name,
		Status: "active",
	}
	
	course := Repository.CreateCourse(c, courseCreate)
	if(course.ID == 0) {
		c.JSON(http.StatusOK, gin.H{"status" : false, "message": "Duplicate Name" })
		return
	}
	c.JSON(http.StatusOK, gin.H{"status" : true, "message": "OK" , "data":course})
}