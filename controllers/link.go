package controllers

import (
	"clincker/interfaces"
	"clincker/models"
	"clincker/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type LinkController struct {
	ListByUser func(request *gin.Context)
	Create     func(request *gin.Context)
	Update     func(request *gin.Context)
	Delete     func(request *gin.Context)
	Show       func(request *gin.Context)
}

func Link() LinkController {
	return LinkController{
		ListByUser: listByUserLink,
		Create:     createLink,
		Update:     updateLink,
		Delete:     deleteLink,
		Show:       showLink,
	}
}

func listByUserLink(request *gin.Context) {
	type response struct {
		Resource interfaces.Response `json:"resource"`
		Data     []models.LinkStruct `json:"links"`
	}

	userId := request.Param("user_id")
	userIdFormat, _ := strconv.Atoi(userId)

	user, userError := models.User().Show(userIdFormat)

	if userError != nil {
		messageContent := userError.Error()

		if utils.Log().IsNoRowsError(messageContent) {
			messageContent = fmt.Sprintf(
				"Usuário %d não encontrado.",
				userIdFormat)
		}

		request.IndentedJSON(http.StatusBadRequest, interfaces.Response{
			Ok:      false,
			Message: messageContent,
		})

		return
	}

	links, linksException := models.Link().ListByUser(user.Id)

	if linksException != nil {
		if utils.Log().IsNoRowsError(linksException.Error()) {
			request.IndentedJSON(http.StatusOK, response{
				Resource: interfaces.Response{
					Ok: true,
					Message: fmt.Sprintf(
						"Links do usuário %d.", userIdFormat),
				},
				Data: links,
			})

			return
		}

		request.IndentedJSON(http.StatusBadRequest, interfaces.Response{
			Ok: false,
			Message: fmt.Sprintf(
				"Erro ao buscar os links do usuário %d.", user.Id),
		})

		return
	}

	request.IndentedJSON(http.StatusOK, response{
		Resource: interfaces.Response{
			Ok: true,
			Message: fmt.Sprintf(
				"Links do usuário %d.", userIdFormat),
		},
		Data: links,
	})
}

func createLink(request *gin.Context) {
	type response struct {
		Resource interfaces.Response          `json:"resource"`
		Data     models.NewLinkResponseStruct `json:"link"`
	}

	var newUser models.UserInsertStruct

	requestError := request.BindJSON(&newUser)

	if requestError != nil {
		request.IndentedJSON(http.StatusOK, interfaces.Response{
			Ok: false,
			Message: fmt.Sprintf(
				"Formato inválido de JSON enviado.",
			),
		})

		return
	}

	if !models.User().IsValid(newUser) {
		request.IndentedJSON(http.StatusOK, interfaces.Response{
			Ok: false,
			Message: fmt.Sprintf(
				"Campos obrigatórios faltando no JSON enviado.",
			),
		})

		return
	}

	userExists, exception := models.User().Verify(newUser.Email)

	if userExists != nil {
		message := ""

		if exception != nil && exception.Error() != "" {
			message = exception.Error()
		}

		request.IndentedJSON(http.StatusOK, response{
			Resource: interfaces.Response{
				Ok: false,
				Message: fmt.Sprintf(
					"Erro na inserção de usuário! [%s]",
					message,
				),
			},
			Data: models.NewUserResponseStruct{
				Id: userExists.Id,
			},
		})

		return
	}

	hashedPassword, exception := utils.Crypto().Hash(newUser.Password)

	if exception != nil {
		request.IndentedJSON(http.StatusOK, interfaces.Response{
			Ok: false,
			Message: fmt.Sprintf(
				"Erro na inserção de usuário! [%s]", exception.Error(),
			),
		})

		return
	}

	newUser.Password = hashedPassword

	id, exception := models.User().Create(newUser)

	if exception != nil {
		request.IndentedJSON(http.StatusOK, interfaces.Response{
			Ok: false,
			Message: fmt.Sprintf(
				"Erro na inserção de usuário! [%s]", exception.Error(),
			),
		})

		return
	}

	request.IndentedJSON(http.StatusOK, response{
		Resource: interfaces.Response{
			Ok: true,
			Message: fmt.Sprintf(
				"Rota de Inserção de Usuário com código %d!", id),
		},
		Data: models.NewUserResponseStruct{
			Id: id,
		},
	})
}

func updateLink(request *gin.Context) {}

func deleteLink(request *gin.Context) {}

func showLink(request *gin.Context) {}
