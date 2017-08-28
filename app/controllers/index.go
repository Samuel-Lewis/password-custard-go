package controllers

import (
	"html/template"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("app/views/index.html")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	pack := "hello!"

	err = tmpl.Execute(w, pack)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
