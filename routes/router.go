package routes

import (
	"clincker/controllers"
	"github.com/gin-gonic/gin"
	"log"
)

var router *gin.Engine

func Start() {
	boot()
	setup()
	listen()
}

func boot() {
	router = gin.Default()
}

func setup() {
	router.GET("/", controllers.Hello)
}

func listen() {
	requestError := router.Run(":8080")

	if requestError != nil {
		log.Fatal(requestError)
	}
}
