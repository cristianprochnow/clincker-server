package controllers

import (
	"clincker/interfaces"
	"clincker/models"
	"clincker/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Login  func(request *gin.Context)
	List   func(request *gin.Context)
	Show   func(request *gin.Context)
	Create func(request *gin.Context)
	Update func(request *gin.Context)
	Delete func(request *gin.Context)
}

func User() UserController {
	return UserController{
		Login:  login,
		List:   list,
		Show:   show,
		Create: create,
		Update: update,
		Delete: remove,
	}
}

func login(request *gin.Context) {
	type response struct {
		Resource interfaces.Response   `json:"resource"`
		Data     models.UserLoginToken `json:"login"`
	}

	var user models.UserLogin

	requestError := request.BindJSON(&user)

	if requestError != nil {
		request.IndentedJSON(http.StatusOK, interfaces.Response{
			Ok: false,
			Message: fmt.Sprintf(
				"Formato inválido de JSON enviado.",
			),
		})

		return
	}

	if !models.User().IsValidLogin(user) {
		request.IndentedJSON(http.StatusOK, interfaces.Response{
			Ok: false,
			Message: fmt.Sprintf(
				"Campos obrigatórios faltando no JSON enviado.",
			),
		})

		return
	}

	userExists, _ := models.User().Verify(user.Email)

	if userExists == nil {
		request.IndentedJSON(http.StatusOK, interfaces.Response{
			Ok: false,
			Message: fmt.Sprintf(
				"Usuário não encontrado no banco de dados.",
			),
		})

		return
	}

	matches := utils.Crypto().Equals(userExists.Password, user.Password)

	if !matches {
		request.IndentedJSON(http.StatusOK, interfaces.Response{
			Ok:      false,
			Message: fmt.Sprintf("Usuário ou senha inválidos"),
		})

		return
	}

	userHash, _ := models.User().GetHash(userExists.Id)

	token, tokenError := utils.Crypto().Hash(
		utils.User().GetLoginToken(userHash),
	)

	if tokenError != nil {
		request.IndentedJSON(http.StatusOK, interfaces.Response{
			Ok: false,
			Message: fmt.Sprintf(
				"Erro ao gerar o token de validação [%s]",
				tokenError.Error()),
		})

		return
	}

	request.IndentedJSON(http.StatusOK, response{
		Resource: interfaces.Response{
			Ok:      true,
			Message: "Rota de login do usuário!",
		},
		Data: models.UserLoginToken{
			Id:    userExists.Id,
			Token: token,
		},
	})
}

func list(request *gin.Context) {
	users, exception := models.User().List()

	if exception != nil {
		request.IndentedJSON(http.StatusOK, interfaces.Response{
			Ok: false,
			Message: fmt.Sprintf(
				"Erro na listagem de usuários! [%s]", exception.Error(),
			),
		})

		return
	}

	type response struct {
		Resource interfaces.Response `json:"resource"`
		Data     []models.UserStruct `json:"users"`
	}

	request.IndentedJSON(http.StatusOK, response{
		Resource: interfaces.Response{
			Ok:      exception == nil,
			Message: "Rota de listagem de Usuários!",
		},
		Data: users,
	})
}

func show(request *gin.Context) {
	type response struct {
		Resource interfaces.Response `json:"resource"`
		Data     *models.UserStruct  `json:"user"`
	}

	id := request.Param("id")
	value, conversionError := strconv.Atoi(id)

	if conversionError != nil {
		request.IndentedJSON(http.StatusOK, interfaces.Response{
			Ok:      false,
			Message: fmt.Sprintf("Valor de ID %s inválido!", id),
		})

		return
	}

	user, exception := models.User().Show(value)

	if exception != nil {
		if utils.Log().IsNoRowsError(exception.Error()) {
			request.IndentedJSON(http.StatusOK, response{
				Resource: interfaces.Response{
					Ok:      true,
					Message: fmt.Sprintf("Rota de Info de Usuário com código %s!", id),
				},
				Data: nil,
			})

			return
		}

		request.IndentedJSON(http.StatusOK, interfaces.Response{
			Ok: false,
			Message: fmt.Sprintf(
				"Erro na consulta de usuário! [%s]", exception.Error(),
			),
		})

		return
	}

	request.IndentedJSON(http.StatusOK, response{
		Resource: interfaces.Response{
			Ok:      true,
			Message: fmt.Sprintf("Rota de Info de Usuário com código %s!", id),
		},
		Data: user,
	})
}

func create(request *gin.Context) {
	type response struct {
		Resource interfaces.Response          `json:"resource"`
		Data     models.NewUserResponseStruct `json:"user"`
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

func update(request *gin.Context) {
	type response struct {
		Resource interfaces.Response          `json:"resource"`
		Data     models.NewUserResponseStruct `json:"user"`
	}

	id := request.Param("id")
	value, conversionError := strconv.Atoi(id)

	if conversionError != nil {
		request.IndentedJSON(http.StatusOK, interfaces.Response{
			Ok:      false,
			Message: fmt.Sprintf("Valor de ID %s inválido!", id),
		})

		return
	}

	verificationUser, exception := models.User().Show(value)

	if verificationUser == nil {
		request.IndentedJSON(http.StatusOK, interfaces.Response{
			Ok: false,
			Message: fmt.Sprintf(
				"Usuário %d não encontrado na base de dados", value,
			),
		})

		return
	}

	var userExists models.UserInsertStruct

	requestError := request.BindJSON(&userExists)

	if requestError != nil {
		request.IndentedJSON(http.StatusOK, interfaces.Response{
			Ok: false,
			Message: fmt.Sprintf(
				"Formato inválido de JSON enviado.",
			),
		})

		return
	}

	if !models.User().IsValid(userExists) {
		request.IndentedJSON(http.StatusOK, interfaces.Response{
			Ok: false,
			Message: fmt.Sprintf(
				"Campos obrigatórios faltando no JSON enviado.",
			),
		})

		return
	}

	otherUser, _ := models.User().Verify(userExists.Email)

	if otherUser != nil &&
		otherUser.Email == userExists.Email &&
		otherUser.Id != value {
		request.IndentedJSON(http.StatusOK, response{
			Resource: interfaces.Response{
				Ok: false,
				Message: fmt.Sprintf(
					"Erro na inserção de usuário! "+
						"[Usuário com e-mail %s já existe com id %d!]",
					userExists.Email, otherUser.Id),
			},
			Data: models.NewUserResponseStruct{
				Id: verificationUser.Id,
			},
		})

		return
	}

	hashedPassword, exception := utils.Crypto().Hash(userExists.Password)

	if exception != nil {
		request.IndentedJSON(http.StatusOK, interfaces.Response{
			Ok: false,
			Message: fmt.Sprintf(
				"Erro na inserção de usuário! [%s]", exception.Error(),
			),
		})

		return
	}

	userExists.Password = hashedPassword

	_, exception = models.User().Update(userExists, value)

	if exception != nil {
		request.IndentedJSON(http.StatusOK, interfaces.Response{
			Ok: false,
			Message: fmt.Sprintf(
				"Erro na Atualização de usuário! [%s]", exception.Error(),
			),
		})

		return
	}

	request.IndentedJSON(http.StatusOK, response{
		Resource: interfaces.Response{
			Ok: true,
			Message: fmt.Sprintf(
				"Rota de Atualização de Usuário com código %d!", value),
		},
		Data: models.NewUserResponseStruct{
			Id: value,
		},
	})
}

func remove(request *gin.Context) {
	id := request.Param("id")
	value, conversionError := strconv.Atoi(id)

	if conversionError != nil {
		request.IndentedJSON(http.StatusOK, interfaces.Response{
			Ok:      false,
			Message: fmt.Sprintf("Valor de ID %s inválido!", id),
		})

		return
	}

	verificationUser, exception := models.User().Show(value)

	if verificationUser == nil {
		request.IndentedJSON(http.StatusOK, interfaces.Response{
			Ok: false,
			Message: fmt.Sprintf(
				"Usuário %d não encontrado na base de dados", value,
			),
		})

		return
	}

	_, exception = models.User().Delete(value)

	if exception != nil {
		request.IndentedJSON(http.StatusOK, interfaces.Response{
			Ok: false,
			Message: fmt.Sprintf(
				"Erro na Exclusão de usuário! [%s]", exception.Error(),
			),
		})

		return
	}

	request.IndentedJSON(http.StatusOK, interfaces.Response{
		Ok: true,
		Message: fmt.Sprintf(
			"Usuário excluído com sucesso! [%d]", value,
		),
	})
}
