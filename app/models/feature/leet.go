package feature

import (
	"github.com/Samuel-Lewis/Password-Custard/app/models"
)

type location struct {
	word int
	pos  int
}

// Leet replaces letters with common symbols or numbers
func Leet(s []string) []string {

	points := make([]location, 0)
	// Find locations of all replaceable chars
	for i, w := range s {
		for j, c := range w {
			if models.HasLeet(byte(c)) {
				points = append(points, location{word: i, pos: j})
			}
		}
	}

	if len(points) == 0 {
		return s
	}

	// Replace the char
	p := points[models.GetRand(0, len(points))]
	w := []byte(s[p.word])
	w[p.pos] = models.GetLeet(w[p.pos])
	s[p.word] = string(w)
	return s
}
