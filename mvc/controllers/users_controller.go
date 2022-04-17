package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/adityasunny1189/REST-Microservice-Golang/mvc/services"
)

func GetUsers(res http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("user id must be a number"))
		return
	}

	user, err := services.GetUser(userId)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte(err.Error()))
		return
	}

	jsondata, _ := json.Marshal(user)
	res.Write(jsondata)
}
