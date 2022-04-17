package app

import (
	"net/http"

	"github.com/adityasunny1189/REST-Microservice-Golang/mvc/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUsers)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
