package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	//carregando as variaveis de ambiente com godotenv
	errGoDotEnv := godotenv.Load(".env")
	//validando erros
	if errGoDotEnv != nil {
		log.Fatal(errGoDotEnv)
	}

	fmt.Println(os.Getenv("TEST"))
}
