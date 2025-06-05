package main

import "go.uber.org/zap"

func main() {

	//pe possivel configurar seu proprip logger semsare o NewProducion

	/*

		loggerConfig := zap.Config{
			//as colunas de ocnfiguração
		}

		myLogger := zap.Builder(loggerConfig) -> retorna um *zap.Logger

	*/

	//inicializar o logger
	logger, _ := zap.NewProduction() //ele retorna dosi valores, o looger e um erro

}
