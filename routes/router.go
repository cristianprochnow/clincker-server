package routes

import (
	"clincker/controllers"
	"clincker/middlewares"
	"clincker/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Start() {
	boot()
	setup()
	listen()
}

func boot() {
	router = gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{
		"Origin",
		"Content-Length",
		"Content-Type",
		"CLINCKER-TOKEN",
		"CLINCKER-USER",
	}

	router.Use(cors.New(config))
}

func setup() {
	router.GET(
		"/",
		controllers.Hello().Hello)

	userRoutes := router.Group("/user")
	{
		userRoutes.GET(
			"/",
			middlewares.Auth().Verify,
			controllers.User().List)
		userRoutes.GET(
			"/:id",
			middlewares.Auth().Verify,
			controllers.User().Show)
		userRoutes.POST("/", controllers.User().Create)
		userRoutes.PUT(
			"/:id",
			middlewares.Auth().Verify,
			controllers.User().Update)
		userRoutes.DELETE(
			"/:id",
			middlewares.Auth().Verify,
			controllers.User().Delete)
		userRoutes.POST("/login", controllers.User().Login)
	}

	linkRoutes := router.Group("/link")
	{
		linkRoutes.GET(
			"/user/:user_id",
			middlewares.Auth().Verify,
			controllers.Link().ListByUser)
		linkRoutes.POST(
			"/",
			middlewares.Auth().Verify,
			controllers.Link().Create)
		linkRoutes.PUT(
			"/:link_id",
			middlewares.Auth().Verify,
			controllers.Link().Update)
		linkRoutes.GET(
			"/:link_id",
			middlewares.Auth().Verify,
			controllers.Link().Show)
		linkRoutes.DELETE(
			"/:link_id",
			middlewares.Auth().Verify,
			controllers.Link().Delete)
	}
}

func listen() {
	utils.Log().Exception(router.Run(":8080"))
}
