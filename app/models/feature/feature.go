package feature

import (
	"log"
	"math/rand"
)

// Applier is function type for function to take the current password, manipulate, and return it
type Applier func([]string) []string

var features map[string]Applier

// Nop is the default and performs no operation to the password
func Nop(s []string) []string {
	log.Printf("WARNING! Using Nop feature. This maybe from using a wrong key to a feature")
	return s
}

// Register builds the feature map from string to feature implementations
func Register() {
	features = make(map[string]Applier)

	features["noun"] = Noun
	features["verb"] = Verb
	features["adjective"] = Adjective

	log.Printf("Registered %d features", len(features))
}

// Get finds requested feature and returns t
func Get(s string) Applier {
	if val, ok := features[s]; ok {
		return val
	}
	return Nop
}

// Insert value at random position in password
func Insert(s string, p []string) []string {
	if len(p) == 0 {
		log.Printf("len(p) == 0")
		return []string{s}
	}

	i := rand.Intn(len(p))
	p = append(p, "")
	copy(p[i+1:], p[i:])
	p[i] = s
	return p
}
