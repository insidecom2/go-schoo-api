package main

import (
	"go/go-api/model"
	"go/go-api/route"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func main() {
	model.InitModel()
	
	r := gin.Default()
	r.Use(cors.Default())
	/* Routes*/
	route.AuthRoutes()

	r.Run("localhost:9999") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}


