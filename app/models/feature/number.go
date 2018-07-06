package feature

import (
	"fmt"
	"math/rand"
)

// NumberSimple inserts a random number between words
func NumberSimple(s []string) []string {
	n := rand.Intn(1000)
	return Insert(fmt.Sprint(n), s)
}

// NumberRandom inserts a random Number into an exsiting word
func NumberRandom(s []string) []string {
	// TODO
	return NumberSimple(s)
}
