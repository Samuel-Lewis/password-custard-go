package feature

import (
	"strings"

	"github.com/Samuel-Lewis/Password-Custard/app/models"
)

// TitleCase changes all words to title case
func TitleCase(s []string) []string {
	for i := 0; i < len(s); i++ {
		s[i] = strings.Title(s[i])
	}
	return s
}

// UpperCase transforms one word to all upper case
func UpperCase(s []string) []string {
	l := len(s)
	if l == 0 {
		return s
	}

	i := models.GetRand(0, l)
	s[i] = strings.ToUpper(s[i])

	// Make following word lower, for readability
	if i < l-1 {
		s[i+1] = strings.ToLower(s[i+1])
	}

	return s
}
