package route

import (
	AuthController "go/go-api/controller/auth"
	CourseController "go/go-api/controller/course"
	UserController "go/go-api/controller/user"
	Middleware "go/go-api/middleware"

	"github.com/gin-gonic/gin"
)


func AuthRoutes (r *gin.Engine) {
	r.POST("/register", AuthController.Register)
	r.POST("/login",AuthController.Login)
}

func UserRoutes(r *gin.Engine) {
	authorized := r.Group("/user",Middleware.AuthMiddleware)
	authorized.GET("/",UserController.GetProfile)
}

func CourseRoutes(r *gin.Engine) {
	authorized := r.Group("/course",Middleware.AuthMiddleware)
	authorized.POST("/",CourseController.CreateCourse)
}