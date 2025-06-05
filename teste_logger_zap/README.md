# Logger usando zap

usando o método ``` zap.NewProduction() ``` ele cria o logger

## forma de criar um logger personalizado
```
		loggerConfig := zap.Config{
			//as colunas de ocnfiguração
		}

		myLogger := zap.Builder(loggerConfig) -> retorna um *zap.Logger

```