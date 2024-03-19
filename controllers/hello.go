package controllers

import (
	"clincker/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(request *gin.Context) {
	request.IndentedJSON(http.StatusOK, models.Response{
		Ok:      true,
		Message: "Clincker!",
	})
}
