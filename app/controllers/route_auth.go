package controllers

import (
	"app/models"
	"net/http"
	"html/template"
	"log"
	"context"
)

func signup(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		renderSignupTemplate(w)
	case http.MethodPost:
		processSignupForm(w, r)
	default:
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}

func renderSignupTemplate(w http.ResponseWriter) {
	t, err := template.ParseFiles("views/templates/signup.html")
	if err != nil {
		log.Fatalln("Template parsing error:", err)
	}
	t.Execute(w, nil)
}

func processSignupForm(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	if err := r.ParseForm(); err != nil {
		log.Fatalln("Form parsing error:", err)
	}
	user := models.User{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	if err := user.CreateUser(ctx); err != nil {
		log.Fatalln("User creation error:", err)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}