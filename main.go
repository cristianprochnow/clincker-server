package main

import (
	"clincker/controllers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", controllers.Hello)

	requestError := router.Run(":8080")

	if requestError != nil {
		log.Fatal(requestError)
	}
}
