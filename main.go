package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/Samuel-Lewis/PassGen/app/controllers"
	"github.com/Samuel-Lewis/PassGen/app/models"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	rand.Seed(time.Now().Unix())

	models.ReadWords("db/verbs.txt", models.VerbList)
	models.ReadWords("db/adjectives.txt", models.AdjList)
	models.ReadWords("db/nouns.txt", models.NounList)
	models.ReadLeet("db/leet.txt")
	models.ReadSymbols("db/symbols.txt")

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", controllers.Index)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
