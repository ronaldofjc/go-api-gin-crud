package main

import (
	"api-crud/src/configuration/logger"
	"api-crud/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	logger.Info("About to start user application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//envTest := os.Getenv("HELLO")

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8090"); err != nil {
		log.Fatal(err)
	}
}
