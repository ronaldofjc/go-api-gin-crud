package main

import (
	"api-crud/src/configuration/database/mongodb"
	"api-crud/src/configuration/logger"
	"api-crud/src/controller"
	"api-crud/src/controller/routes"
	"api-crud/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	logger.Info("Start user application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongodb.InitConnection()

	//Init dependencies
	userService := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(userService)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8090"); err != nil {
		log.Fatal(err)
	}
}
