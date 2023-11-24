package controllers

import (
	"app/models"
	"net/http"
	"html/template"
	"log"
	"context"
)

func signup(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err == nil {
		http.Redirect(w, r, "/todos", http.StatusFound)
	}

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
	t, err := template.ParseFiles("views/templates/signup.html", "views/templates/navbar.html")
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

func login(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err == nil {
		http.Redirect(w, r, "/todos", http.StatusFound)
	}

	t, err := template.ParseFiles("views/templates/login.html", "views/templates/navbar.html")
	if err != nil {
		log.Fatalln("Template parsing error:", err)
	}
	t.Execute(w, nil)
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err == nil {
		http.Redirect(w, r, "/todos", http.StatusFound)
	}

	ctx := context.Background()
	if err := r.ParseForm(); err != nil {
		log.Fatalln("Form parsing error:", err)
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	user, err := models.GetUserByEmail(ctx, r.FormValue("email"))
	if err != nil {
		log.Fatalln("User does not exist:", err)
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	if user.Password != models.Encrypt(r.FormValue("password")) {
		log.Fatalln("Password is not correct:", err)
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	session, err := user.CreateSession(ctx)
	if err != nil {
		log.Fatalln("Session creation error:", err)
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	cookie := http.Cookie{
		Name:     "_cookie",
		Value:    session.UUID,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/todos", http.StatusFound)
}