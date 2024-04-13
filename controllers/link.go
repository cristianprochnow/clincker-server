package controllers

import (
	"clincker/interfaces"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type LinkController struct {
	ListByUser func(request *gin.Context)
}

func Link() LinkController {
	return LinkController{
		ListByUser: listByUserLink,
	}
}

func listByUserLink(request *gin.Context) {
	userId := request.Param("user_id")
	userIdFormat, _ := strconv.Atoi(userId)

	request.IndentedJSON(http.StatusOK, interfaces.Response{
		Ok: true,
		Message: fmt.Sprintf(
			"Rota de listagem de links do usuário %d!", userIdFormat),
	})
}
