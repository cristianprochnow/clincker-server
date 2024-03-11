package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type response struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func main() {
	router := gin.Default()

	router.GET("/", hello)

	requestError := router.Run(":8080")

	if requestError != nil {
		log.Fatal(requestError)
	}
}

func hello(request *gin.Context) {
	request.IndentedJSON(http.StatusOK, response{
		Ok:      true,
		Message: "Clincker!",
	})
}
