package controllers

import (
	"net/http"
	"app/config"
)

func StartMainServer() error {
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	return http.ListenAndServe(":" + config.Config.AppPort, nil)
}