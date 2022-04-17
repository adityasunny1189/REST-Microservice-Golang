package domain

import "errors"

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

func GetUser(userid int64) (*User, error) {
	if user := users[userid]; user != nil {
		return user, nil
	}
	return nil, errors.New("user not found")
}
