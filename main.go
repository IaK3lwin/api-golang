package main

import (
	"api-mvp/src/controller/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	//carregando as variaveis de ambiente com godotenv
	errGoDotEnv := godotenv.Load(".env")
	//validando erros
	if errGoDotEnv != nil {
		log.Fatal(errGoDotEnv)
	}

	//criar o nosso router
	router := gin.Default()

	//o noso iniRoutem espera a alocação de memória onde nosso router(com rodos os mettodos struct do
	// router group, lembre0se, o init router espera um endereço de memodria de um *gin.RouterGroup )
	routes.InitRoutes(&router.RouterGroup)

	//rodar o servidor de fato
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
