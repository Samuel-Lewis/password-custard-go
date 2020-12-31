package controllers

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/Samuel-Lewis/password-custard-go/app/models"
	"github.com/Samuel-Lewis/password-custard-go/app/models/feature"
)

// PassOut response type to the write
type PassOut struct {
	Password string
}

// Raw handles the /raw call (used for all password generating)
func Raw(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	// TODO: replace with some preset or default?
	q := "words:2:3,numbers,symbols:1:2,uppercase:0:1,titlecase"
	if val, ok := r.URL.Query()["q"]; ok {
		if matched, _ := regexp.MatchString("^([a-z]+(:[0-9]+){0,2},)*([a-z]+(:[0-9]+){0,2}),?$", val[0]); matched {
			q = string(val[0])
		} else {
			w.Write([]byte("[ERROR: Format syntax incorrect]"))
			return
		}
	}
	q = strings.TrimSuffix(q, ",")
	w.Write([]byte(generatePassword(q)))
}

// GeneratePassword makes the password
func generatePassword(q string) string {
	var feats []string

	items := strings.Split(q, ",")
	for _, i := range items {
		tokens := strings.Split(i, ":")
		s := 1
		e := 1

		if (len(tokens)) > 1 {
			s, _ = strconv.Atoi(tokens[1])
			e = s
		}

		if len(tokens) > 2 {
			e, _ = strconv.Atoi(tokens[2])
		}

		r := models.GetRand(s, e+1)
		for r > 0 {
			r--
			feats = append(feats, feature.Choose(tokens[0]))
		}
	}

	// Orders application of features
	feats = feature.Order(feats)

	// Apply features
	phrase := []string{}
	for _, f := range feats {
		phrase = feature.Apply(f)(phrase)
	}

	return strings.Join(phrase, "")
}
