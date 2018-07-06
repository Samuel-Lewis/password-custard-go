package feature

import (
	"log"
	"math/rand"
	"sort"
)

// Applier is function type for function to take the current password, manipulate, and return it
type Applier func([]string) []string

var features map[string]Applier
var groups map[string][]string
var order []string

// Nop is the default and performs no operation to the password
func Nop(s []string) []string {
	log.Printf("WARNING! Using Nop feature. This maybe from using a wrong key to a feature")
	return s
}

// Register builds the feature map from string to feature implementations
func Register() {
	groups = make(map[string][]string)
	features = make(map[string]Applier)

	groups["word"] = []string{"noun", "verb", "adjective"}
	groups["symbol"] = []string{"symbolsimple", "symbolrandom", "symbolreplace"}
	groups["capitalise"] = []string{"titlecase", "uppercase"}

	features["noun"] = Noun
	features["verb"] = Verb
	features["adjective"] = Adjective
	features["symbolsimple"] = SymbolSimple
	features["symbolrandom"] = SymbolRandom
	features["symbolreplace"] = SymbolReplace
	features["titlecase"] = TitleCase
	features["uppercase"] = UpperCase

	order = []string{
		"noun", "verb", "adjective",
		"titlecase", "uppercase",
		"leet",
		"numberreplace", "symbolreplace",
		"numberrandom", "symbolrandom",
		"numbersimple", "symbolsimple",
	}

	log.Printf("Registered %d features", len(features))
}

// Choose selects a random from a group, or if not a group, returns it self
func Choose(s string) string {
	if val, ok := groups[s]; ok {
		return val[rand.Intn(len(val))]
	}
	return s
}

// Apply finds requested feature and returns t
func Apply(s string) Applier {
	if val, ok := features[s]; ok {
		return val
	}
	return Nop
}

// Insert value at random position in password
func Insert(s string, p []string) []string {
	if len(p) == 0 {
		return []string{s}
	}

	i := rand.Intn(len(p) + 1)
	p = append(p, "")
	copy(p[i+1:], p[i:])
	p[i] = s
	return p
}

type byOrder []string

func (s byOrder) Len() int {
	return len(s)
}

func (s byOrder) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byOrder) Less(i, j int) bool {
	// true if s[i] should come before s[j]
	if s[i] == s[j] {
		return true
	}

	// Check compared to list
	for _, val := range order {
		if val == s[i] {
			return true
		}

		if val == s[j] {
			return false
		}
	}
	return true
}

// Order sorts the methods to apply according to topo order
func Order(s []string) []string {
	sort.Sort(byOrder(s))
	return s
}
