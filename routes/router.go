package routes

import (
	"clincker/controllers"
	"clincker/utils"
	"github.com/gin-gonic/gin"
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
	utils.Error(router.Run(":8080"))
}
