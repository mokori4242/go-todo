package controllers

import (
	"net/http"
	"app/config"
)

func StartMainServer() error {
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	return http.ListenAndServe(":" + config.Config.AppPort, nil)
}