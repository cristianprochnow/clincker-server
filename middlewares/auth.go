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

	if (userError != nil) {
		request.IndentedJSON(http.StatusBadRequest, interfaces.Response{
			Ok: false,
			Message: userError.Error(),
		})
		request.Abort()

		return
	}

	if (user == nil) {
		request.IndentedJSON(http.StatusForbidden, interfaces.Response{
			Ok: false,
			Message: fmt.Sprintf(
				"Usuário %d não encontrado.", userIdFormat
			),
		})
		request.Abort()

		return
	}

	isValidCredentials := utils.Crypto().Equals(
		userToken,
		utils.Crypto().Hash(utils.User().GetLoginToken(
			user.Email, user.Name,
		)),
	)

	if (!isValidCredentials) {
		request.IndentedJSON(http.StatusForbidden, interfaces.Response{
			Ok: false,
			Message: "Token CLINCKER-TOKEN é inválido.",
		})
		request.Abort()

		return
	}

	request.Next()
}
