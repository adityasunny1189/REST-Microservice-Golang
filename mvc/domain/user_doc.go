package domain

import (
	"github.com/adityasunny1189/REST-Microservice-Golang/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {
			ID:        123,
			FirstName: "Aditya",
			LastName:  "Pathak",
			Email:     "adityapathak1189@gmail.com",
		},
	}
)

func GetUser(userid int64) (*User, *utils.ApplicationError) {
	if user := users[userid]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    "user not found",
		StatusCode: 404,
		Code:       "please check user id",
	}
}
