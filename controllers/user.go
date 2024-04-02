package controllers

import (
	"clincker/interfaces"
	"clincker/models"
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
	users, exception := models.User().List()

	type response struct {
		Resource interfaces.Response `json:"resource"`
		Data     []models.UserStruct `json:"users"`
	}

	message := "Listagem de usuários concluída."

	if exception != nil {
		message = exception.Error()
	}

	request.IndentedJSON(http.StatusOK, response{
		Resource: interfaces.Response{
			Ok:      exception == nil,
			Message: message,
		},
		Data: users,
	})
}

func show(request *gin.Context) {
	id := request.Param("id")

	request.IndentedJSON(http.StatusOK, interfaces.Response{
		Ok:      true,
		Message: fmt.Sprintf("Rota de Info de Usuário com código %s!", id),
	})
}
