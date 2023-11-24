package controllers

import (
	"net/http"
	"fmt"
	"app/models"
	"app/config"
	"context"
)

func session(w http.ResponseWriter, r *http.Request) (s models.Session, err error) {
	ctx := context.Background()
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		s = models.Session{UUID: cookie.Value}
		if ok, _ := s.CheckSession(ctx); !ok {
			err = fmt.Errorf("Invalid session")
		}
	}
	return s, err
}

func StartMainServer() error {
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/todos", todos)
	return http.ListenAndServe(":" + config.Config.AppPort, nil)
}