package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/Samuel-Lewis/Password-Custard/app/controllers"
	"github.com/Samuel-Lewis/Password-Custard/app/models"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	rand.Seed(time.Now().Unix())

	models.VerbList = models.ReadWords("db/verbs.txt")
	models.AdjList = models.ReadWords("db/adjectives.txt")
	models.NounList = models.ReadWords("db/nouns.txt")
	models.ReadLeet("db/leet.txt")
	models.ReadSymbols("db/symbols.txt")

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/raw", controllers.Raw)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
