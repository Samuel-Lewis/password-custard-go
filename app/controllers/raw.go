package controllers

import (
	"net/http"
	"strings"

	"github.com/Samuel-Lewis/Password-Custard/app/models/feature"
)

// PassOut response type to the write
type PassOut struct {
	Password string
}

// Raw handles the /raw call (used for all password generating)
func Raw(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(generatePassword()))
}

// GeneratePassword makes the password
func generatePassword() string {

	// Polled features
	var feats []string
	feats = append(feats, feature.Choose("word"))
	feats = append(feats, feature.Choose("word"))
	feats = append(feats, feature.Choose("word"))
	feats = append(feats, feature.Choose("symbol"))
	feats = append(feats, feature.Choose("uppercase"))
	feats = append(feats, feature.Choose("titlecase"))

	// Orders application of features
	feats = feature.Order(feats)

	// Apply features
	phrase := []string{}
	for _, f := range feats {
		phrase = feature.Apply(f)(phrase)
	}

	return strings.Join(phrase, "")
}
