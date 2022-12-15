package routes

import (
	"api-crud/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/user_id/:userId", controller.FindUserById)
	r.GET("/user_email/:userEmail", controller.FindUserByEmail)
	r.POST("/user", controller.CreateUser)
	r.PUT("/user/:userId", controller.UpdateUser)
	r.DELETE("/user/:userId", controller.DeleteUser)
}
