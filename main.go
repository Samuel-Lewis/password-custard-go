package main

import (
	"math/rand"
	"os"
	"time"

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
}
