package models

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// VerbList stored verbs from file
var VerbList []string

// AdjList stored Adjectives from file
var AdjList []string

// NounList stored nouns from file
var NounList []string

// LeetSpeak for random substitutes
var LeetSpeak map[byte][]byte

// Symbols for random insertion
var Symbols []byte

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

	log.Printf("Loaded %d words from '%s'", len(list), path)
}

func ReadSymbols(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	defer file.Close()

	// Clear the list (to not double up)
	Symbols = []byte{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		Symbols = append(Symbols, []byte(scanner.Text())...)
	}

	log.Printf("Loaded %d symbols from '%s'", len(Symbols), path)
}

func ReadLeet(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	defer file.Close()

	// Clear the list (to not double up)
	LeetSpeak = map[byte][]byte{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var key byte
		var datas string
		fmt.Sscanf(scanner.Text(), "%c | %s", &key, &datas)
		LeetSpeak[key] = []byte(datas)
	}

	log.Printf("Loaded %d characters from '%s'", len(LeetSpeak), path)
}
