package services

import (
	"github.com/adityasunny1189/REST-Microservice-Golang/mvc/domain"
)

func GetUser(userid int64) (*domain.User, error) {
	return domain.GetUser(userid)
}
