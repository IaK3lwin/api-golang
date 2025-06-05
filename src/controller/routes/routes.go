package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(group *gin.RouterGroup) {
	//criando as rotas

	group.GET("/getUsersById/:userId")
	group.GET("/getUsersByEmail/:userEmail")
	group.POST("/insertUser")
	group.PUT("/updateUser/:userId")
	group.DELETE("/deleteUser/:userId")

}
