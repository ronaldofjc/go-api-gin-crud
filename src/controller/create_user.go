package controller

import (
	"api-crud/src/configuration/validation"
	"api-crud/src/controller/model/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, userRequest)
}
