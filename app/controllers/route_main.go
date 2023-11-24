package controllers

import (
	"net/http"
	"html/template"
	"log"
)

func index(w http.ResponseWriter, r *http.Request) {
	i, err := template.ParseFiles("views/templates/index.html", "views/templates/navbar.html")

	if err != nil {
		log.Fatalln(err)
	}

	i.Execute(w, nil)
}