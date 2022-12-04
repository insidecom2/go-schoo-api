package main

import (
	"fmt"
	"go/go-api/model"
	"go/go-api/route"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Cannot Load ENV")
	}
	
	model.InitModel()
	
	r := gin.Default()
	r.Use(cors.Default())
	/* Routes*/
	route.AuthRoutes(r)
	route.UserRoutes(r)

	port := "localhost:"+os.Getenv("PORT")
	r.Run(port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}


