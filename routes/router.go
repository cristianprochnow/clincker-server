package routes

import (
	"clincker/controllers"
	"clincker/middlewares"
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
	router.GET(
		"/",
		middlewares.CORS().Unleash,
		controllers.Hello().Hello)

	userRoutes := router.Group("/user")
	{
		userRoutes.GET(
			"/",
			middlewares.CORS().Unleash,
			middlewares.Auth().Verify,
			controllers.User().List)
		userRoutes.GET(
			"/:id",
			middlewares.CORS().Unleash,
			middlewares.Auth().Verify,
			controllers.User().Show)
		userRoutes.POST("/", controllers.User().Create)
		userRoutes.PUT(
			"/:id",
			middlewares.CORS().Unleash,
			middlewares.Auth().Verify,
			controllers.User().Update)
		userRoutes.DELETE(
			"/:id",
			middlewares.CORS().Unleash,
			middlewares.Auth().Verify,
			controllers.User().Delete)
		userRoutes.POST("/login", controllers.User().Login)
	}

	linkRoutes := router.Group("/link")
	{
		linkRoutes.GET(
			"/user/:user_id",
			middlewares.CORS().Unleash,
			middlewares.Auth().Verify,
			controllers.Link().ListByUser)
		linkRoutes.POST(
			"/",
			middlewares.CORS().Unleash,
			middlewares.Auth().Verify,
			controllers.Link().Create)
		linkRoutes.PUT(
			"/:link_id",
			middlewares.CORS().Unleash,
			middlewares.Auth().Verify,
			controllers.Link().Update)
		linkRoutes.GET(
			"/:link_id",
			middlewares.CORS().Unleash,
			middlewares.Auth().Verify,
			controllers.Link().Show)
		linkRoutes.DELETE(
			"/:link_id",
			middlewares.CORS().Unleash,
			middlewares.Auth().Verify,
			controllers.Link().Delete)
	}
}

func listen() {
	utils.Log().Exception(router.Run(":8080"))
}
