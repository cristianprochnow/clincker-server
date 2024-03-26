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
	router = gin.New()
}

func setup() {
	router.GET("/", controllers.Hello().Hello)

	userRoutes := router.Group("/user")
	{
		userRoutes.GET("/", controllers.User().List)
		userRoutes.GET("/:id", controllers.User().Show)
	}
}

func listen() {
	requestError := router.Run(":8080")

	if requestError != nil {
		log.Fatal(requestError)
	}
}
