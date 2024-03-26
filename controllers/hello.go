package controllers

import (
	"clincker/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloController struct {
	Hello func(request *gin.Context)
}

func Hello() HelloController {
	return HelloController{
		Hello: hello,
	}
}

func hello(request *gin.Context) {
	request.IndentedJSON(http.StatusOK, interfaces.Response{
		Ok:      true,
		Message: "Clincker!",
	})
}
