package controllers

import (
	"clincker/interfaces"
	"clincker/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

	if exception != nil {
		request.IndentedJSON(http.StatusOK, interfaces.Response{
			Ok: false,
			Message: fmt.Sprintf(
				"Erro na listagem de usuários! [%s]", exception.Error(),
			),
		})

		return
	}

	type response struct {
		Resource interfaces.Response `json:"resource"`
		Data     []models.UserStruct `json:"users"`
	}

	request.IndentedJSON(http.StatusOK, response{
		Resource: interfaces.Response{
			Ok:      exception == nil,
			Message: "Rota de listagem de Usuários!",
		},
		Data: users,
	})
}

func show(request *gin.Context) {
	id := request.Param("id")
	value, conversionError := strconv.Atoi(id)

	if conversionError != nil {
		request.IndentedJSON(http.StatusOK, interfaces.Response{
			Ok:      false,
			Message: fmt.Sprintf("Valor de ID %s inválido!", id),
		})

		return
	}

	user, exception := models.User().Show(value)

	if exception != nil {
		request.IndentedJSON(http.StatusOK, interfaces.Response{
			Ok: false,
			Message: fmt.Sprintf(
				"Erro na consulta de usuário! [%s]", exception.Error(),
			),
		})

		return
	}

	type response struct {
		Resource interfaces.Response `json:"resource"`
		Data     *models.UserStruct  `json:"users"`
	}

	request.IndentedJSON(http.StatusOK, response{
		Resource: interfaces.Response{
			Ok:      true,
			Message: fmt.Sprintf("Rota de Info de Usuário com código %s!", id),
		},
		Data: user,
	})
}

func create(request *gin.Context) {

}
