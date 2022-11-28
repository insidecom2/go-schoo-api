package route

import (
	AuthController "go/go-api/controller/auth"

	"github.com/gin-gonic/gin"
)


func AuthRoutes () {
	r := gin.Default()
	r.POST("/register", AuthController.Register)
}