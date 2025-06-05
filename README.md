# projeto api em go arquitetura mvp

## estrutura de pasta

```sh
    SRC  -> arquivos importantes
    |
    |--- Controllers // Entrada e validação de dados
    |
    |--- model // Regra de negócios e objetos principais
    |
    |--- View // Gerenciamento de dados ( publicos ou privados) 
    Converters
    |--- tester // testes de integrações da aplicação
    |--- configurations // Pacote para conexões e configurações
    |
    etc..
    main.go
    .env
    .gitignore // ignora arquivos sensiveis como o .env e não leva para o repositorio remoto
    .readme.md
```


## rotas com gin e gin group

src
|
|--controller
|----- routes
|---------routes.go
|

<code> 

package routes

import (
	"api-mvp/src/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(group *gin.RouterGroup) {
	//criando as rotas

	group.GET("/getUsersById/:userId", controller.FindUserId)
	group.GET("/getUsersByEmail/:userEmail", controller.FundUserByEmail)
	group.POST("/createUser", controller.CreateUser)
	group.PUT("/updateUser/:userId", controller.UpdateUser)
	group.DELETE("/deleteUser/:userId", controller.DeleteUser)

}

</code>


<section>

    InitRoutes -> resonsável oela inicialização das rotas do conroller, no main.go a gente faz a ligaçao dele
    com o roteador do gin ```  router := gin.Default() ```

</section>


 ## Validando as entradas com go validation

```sh
    Quando usamos o validator com o gin, para fazer a validação usamos o "BINDING" ao invés do "VALIDATE", 
    caso usasse somente o validator em go.

    ex:

    type Teste struct {
        Name string `validator:"<valores>"`
    }

    ## com o gin fica:

    type Teste struct {
        Name string `json:"nome" binding:""`
    }



```