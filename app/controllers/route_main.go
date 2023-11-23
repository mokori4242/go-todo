package controllers

import (
	"net/http"
	"html/template"
	"log"
)

func index(w http.ResponseWriter, r *http.Request) {
	i, err := template.ParseFiles("views/templates/signup.html")

	if err != nil {
		log.Fatalln(err)
	}

	i.Execute(w, nil)
}