package models

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

var types map[string][]string
var leet map[byte][]byte
var symbols []byte

// Preload called on wake, loads data to model
func Preload() {
	types = make(map[string][]string)
	types["verb"] = readWords("db/verbs.txt")
	types["adjective"] = readWords("db/adjectives.txt")
	types["noun"] = readWords("db/nouns.txt")

	leet = readLeet("db/leet.txt")
	symbols = readSymbols("db/symbols.txt")
}

// GetWord returns a random word from a given word type
func GetWord(t string) string {
	if val, ok := types[t]; ok {
		return val[rand.Intn(len(val))]
	}

	log.Printf("WARNING! Tried to access a word type that doesn't exist! '%s'", t)
	return "hunter"
}

func readWords(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err.Error())
		return []string{}
	}
	defer file.Close()

	// Clear the list (to not double up)
	list := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		if len(word) > 6 {
			continue
		}
		list = append(list, word)
	}

	log.Printf("Loaded %d words from '%s'", len(list), path)
	return list
}

// readSymbols reads a list of words and returns it
func readSymbols(path string) []byte {
	s := []byte{}
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err.Error())
		return s
	}
	defer file.Close()

	// Clear the list (to not double up)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s = append(symbols, []byte(scanner.Text())...)
	}

	log.Printf("Loaded %d symbols from '%s'", len(s), path)
	return s
}

// readLeet gets a table of chars corresponding to different symbol substitutions
func readLeet(path string) map[byte][]byte {
	l := map[byte][]byte{}
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err.Error())
		return l
	}
	defer file.Close()

	// Clear the list (to not double up)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var key byte
		var datas string
		fmt.Sscanf(scanner.Text(), "%c | %s", &key, &datas)
		l[key] = []byte(datas)
	}

	log.Printf("Loaded %d characters from '%s'", len(l), path)
	return l
}
