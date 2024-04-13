package middlewares

import (
	"clincker/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

	request.Next()
}
