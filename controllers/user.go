package controllers

import (
	"clincker/interfaces"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	List func(request *gin.Context)
	Show func(request *gin.Context)
}

func User() UserController {
	return UserController{
		List: list,
		Show: show,
	}
}

func list(request *gin.Context) {
	request.IndentedJSON(http.StatusOK, interfaces.Response{
		Ok:      true,
		Message: "Rota de listagem dos usuários!",
	})
}

func show(request *gin.Context) {
	id := request.Param("id")

	request.IndentedJSON(http.StatusOK, interfaces.Response{
		Ok:      true,
		Message: fmt.Sprintf("Rota de Info de Usuário com código %s!", id),
	})
}
