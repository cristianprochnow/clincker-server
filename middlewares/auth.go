package middlewares

import (
	"clincker/interfaces"
	"clincker/models"
	"clincker/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	Verify func(request *gin.Context)
}

func Auth() AuthMiddleware {
	return AuthMiddleware{
		Verify: verifyAuth,
	}
}

func verifyAuth(request *gin.Context) {
	userId := request.GetHeader("CLINCKER-USER")
	userIdFormat, _ := strconv.Atoi(userId)
	userToken := request.GetHeader("CLINCKER-TOKEN")

	if userIdFormat == 0 {
		request.IndentedJSON(http.StatusForbidden, interfaces.Response{
			Ok: false,
			Message: "Código de usuário é obrigatório para autenticação " +
				"por meio do Header CLINCKER-USER.",
		})
		request.Abort()

		return
	}

	if userToken == "" {
		request.IndentedJSON(http.StatusForbidden, interfaces.Response{
			Ok: false,
			Message: "Token do usuário é obrigatório para autenticação " +
				"por meio do Header CLINCKER-TOKEN.",
		})
		request.Abort()

		return
	}

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
		request.Abort()

		return
	}

	if userError != nil {
		request.IndentedJSON(http.StatusBadRequest, interfaces.Response{
			Ok:      false,
			Message: userError.Error(),
		})
		request.Abort()

		return
	}

	if user == nil {
		request.IndentedJSON(http.StatusForbidden, interfaces.Response{
			Ok: false,
			Message: fmt.Sprintf(
				"Usuário %d enviado em CLINCKER-USER não foi encontrado.",
				userIdFormat,
			),
		})
		request.Abort()

		return
	}

	hash, _ := utils.Crypto().Hash(utils.User().GetLoginToken(
		user.Email, user.Name,
	))
	isValidCredentials := utils.Crypto().Equals(
		userToken,
		hash,
	)

	if !isValidCredentials {
		request.IndentedJSON(http.StatusForbidden, interfaces.Response{
			Ok:      false,
			Message: "Token CLINCKER-TOKEN é inválido.",
		})
		request.Abort()

		return
	}

	request.Next()
}
