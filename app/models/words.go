package models

import (
	"bufio"
	"log"
	"os"
)

// VerbList stored verbs from file
var VerbList []string

// AdjList stored Adjectives from file
var AdjList []string

// NounList stored nouns from file
var NounList []string

// ReadWords from file
func ReadWords(path string, list []string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	defer file.Close()

	// Clear the list (to not double up)
	list = []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}

	log.Printf("Loaded %d many lines from '%s'", len(list), path)

}
