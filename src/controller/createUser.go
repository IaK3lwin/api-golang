package controller

import (
	errstandard "api-mvp/src/configuration/err_standard"
	"api-mvp/src/controller/model/request"
	"fmt"

	"github.com/gin-gonic/gin"
)

// handler de inserção de usuário
func CreateUser(context *gin.Context) {

	//receber as informações da requiaião
	var userResponse request.UserRequest

	//atribuindo as informações vindas da requisi~ao ao nosso modelo
	if errorShouldByndJson := context.ShouldBindJSON(&userResponse); errorShouldByndJson != nil {

		erro := errstandard.NewErrorBadRequest("parametros de request inválidos")
		fmt.Println(errorShouldByndJson)
		context.JSON(erro.Code, erro)
		return
	}

	context.JSON(200, userResponse)

}
