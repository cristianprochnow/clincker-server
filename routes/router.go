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
		userRoutes.POST("/", controllers.User().Create)
		userRoutes.PUT("/:id", controllers.User().Update)
		userRoutes.POST("/login", controllers.User().Login)
	}

	linkRoutes := router.Group("/link")
	{
		linkRoutes.GET(
			"/user/:user_id",
			controllers.Link().ListByUser)
	}
}

func listen() {
	utils.Log().Exception(router.Run(":8080"))
}
