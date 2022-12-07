package repository

import (
	"go/go-api/model"

	"github.com/gin-gonic/gin"
)
type modelCourse struct {
	ID uint
	Name string 
	Status string 
}

type validateCourse struct {
	Name string 
}

func CreateCourse (c *gin.Context, courseCreate model.Course) modelCourse{
	
	model.Db.Create(&courseCreate) 

	responseCourse := modelCourse{
		ID: courseCreate.ID,
		Name: courseCreate.Name,
		Status: courseCreate.Status,
	}

	return responseCourse
}