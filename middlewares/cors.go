package middlewares

import (
	"github.com/gin-gonic/gin"
)

type CORSMiddleware struct {
	Unleash func(request *gin.Context)
}

func CORS() CORSMiddleware {
	return CORSMiddleware{
		Unleash: unleashCORS,
	}
}

func unleashCORS(request *gin.Context) {
	request.Header("Access-Control-Allow-Origin", "*")
	request.Next()
}
