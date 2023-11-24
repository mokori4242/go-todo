package controllers

import (
	"net/http"
	"html/template"
	"log"
)

func index(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err == nil {
		http.Redirect(w, r, "/todos", http.StatusFound)
	}

	i, err := template.ParseFiles("views/templates/index.html", "views/templates/navbar.html")

	if err != nil {
		log.Fatalln(err)
	}

	i.Execute(w, nil)
}

func todos(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	}

	t, err := template.ParseFiles("views/templates/todos.html", "views/templates/p_navbar.html")
	if err != nil {
		log.Fatalln(err)
	}
	t.Execute(w, nil)
}