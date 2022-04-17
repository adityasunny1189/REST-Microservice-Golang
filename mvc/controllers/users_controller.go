package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/adityasunny1189/REST-Microservice-Golang/mvc/services"
	"github.com/adityasunny1189/REST-Microservice-Golang/mvc/utils"
)

func GetUsers(res http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apperr := &utils.ApplicationError{
			Message:    "user id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "invalid id",
		}
		jsonerrmsg, _ := json.Marshal(apperr)
		res.WriteHeader(apperr.StatusCode)
		res.Write(jsonerrmsg)
		return
	}

	user, apperr := services.GetUser(userId)
	if apperr != nil {
		res.WriteHeader(apperr.StatusCode)
		jsonerrmsg, _ := json.Marshal(apperr)
		res.Write(jsonerrmsg)
		return
	}

	jsondata, _ := json.Marshal(user)
	res.Write(jsondata)
}
