package controllers

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strings"

	"github.com/Samuel-Lewis/Password-Custard/app/models"
)

type PassOut struct {
	Password string
}

func GeneratePassword() string {

	phrase := []string{
		models.VerbList[rand.Intn(len(models.VerbList))],
		models.AdjList[rand.Intn(len(models.AdjList))],
		models.NounList[rand.Intn(len(models.NounList))],
	}

	// Capitalize random word
	// Replace leet
	// Insert numbers
	// Insert symbols

	return strings.Join(phrase, "")
}

func Raw(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("app/views/raw.html")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	pack := PassOut{Password: GeneratePassword()}

	err = tmpl.Execute(w, pack)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
