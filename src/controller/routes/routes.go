package routes

import (
	"api-crud/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {
	r.GET("/user_id/:userId", userController.FindUserByID)
	r.GET("/user_email/:userEmail", userController.FindUserByEmail)
	r.POST("/user", userController.CreateUser)
	r.PUT("/user/:userId", userController.UpdateUser)
	r.DELETE("/user/:userId", userController.DeleteUser)
}
