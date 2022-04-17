package services

import (
	"github.com/adityasunny1189/REST-Microservice-Golang/mvc/domain"
	"github.com/adityasunny1189/REST-Microservice-Golang/mvc/utils"
)

func GetUser(userid int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userid)
}
