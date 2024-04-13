package middlewares

import (
	"fmt"
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
	userToken := request.GetHeader("CLINCKER-TOKEN")

	fmt.Print(userId, userToken)

	request.Next()
}
