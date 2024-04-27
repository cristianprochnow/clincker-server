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
	userId := request.Param("user_id")
	userIdFormat, _ := strconv.Atoi(userId)

	user, userError := models.User().Show(userIdFormat)

	if userError != nil {
		messageContent := userError.Error()

		if utils.Log().IsNoRowsError(messageContent) {
			messageContent = fmt.Sprintf(
				"Usuário %d enviado no CLINCKER-USER não encontrado.",
				userIdFormat)
		}

		request.IndentedJSON(http.StatusBadRequest, interfaces.Response{
			Ok:      false,
			Message: messageContent,
		})

		return
	}

	fmt.Println(user)

	request.IndentedJSON(http.StatusOK, interfaces.Response{
		Ok: true,
		Message: fmt.Sprintf(
			"Rota de listagem de links do usuário %d!", userIdFormat),
	})
}

func createLink(request *gin.Context) {}

func updateLink(request *gin.Context) {}

func deleteLink(request *gin.Context) {}

func showLink(request *gin.Context) {}
