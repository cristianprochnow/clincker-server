package controllers

import (
	"clincker/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	List func(request *gin.Context)
}

func User() UserController {
	return UserController{
		List: list,
	}
}

func list(request *gin.Context) {
	request.IndentedJSON(http.StatusOK, interfaces.Response{
		Ok:      true,
		Message: "Rota de listagem dos usuários!",
	})
}
