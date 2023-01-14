package service

import (
	"api-crud/src/configuration/logger"
	"api-crud/src/configuration/rest_err"
	"api-crud/src/model"
	"fmt"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetPassword())
	return nil
}
