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
