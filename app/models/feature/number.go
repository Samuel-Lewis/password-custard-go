package feature

import (
	"fmt"

	"github.com/Samuel-Lewis/Password-Custard/app/models"
)

// NumberSimple inserts a random number between words
func NumberSimple(s []string) []string {
	n := models.GetRand(0, 1000)
	return Insert(fmt.Sprint(n), s)
}

// NumberRandom inserts a random Number into an exsiting word
func NumberRandom(s []string) []string {
	// TODO
	return NumberSimple(s)
}
