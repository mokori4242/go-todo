package controllers

import (
	"net/http"
	"app/config"
)

func StartMainServer() error {
	http.HandleFunc("/", index)
	return http.ListenAndServe(":" + config.Config.AppPort, nil)
}