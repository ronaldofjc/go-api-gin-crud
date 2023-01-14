package controller

import (
	"api-crud/src/configuration/logger"
	"api-crud/src/configuration/validation"
	"api-crud/src/controller/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller", zap.String("journey", "createUser"))
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "createUser"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	logger.Info("User created successfully", zap.String("journey", "createUser"))
	c.JSON(http.StatusOK, userRequest)
}
