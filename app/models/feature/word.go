package feature

import (
	"github.com/Samuel-Lewis/Password-Custard/app/models"
)

// Noun inserts a random noun into the phrase
func Noun(s []string) []string {
	return Insert(models.GetWord("noun"), s)
}

// Verb inserts a random verb into the phrase
func Verb(s []string) []string {
	return Insert(models.GetWord("verb"), s)
}

// Adjective inserts a random adjective into the phrase
func Adjective(s []string) []string {
	return Insert(models.GetWord("adjective"), s)
}
