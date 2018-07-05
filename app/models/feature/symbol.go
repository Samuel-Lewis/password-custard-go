package feature

import (
	"github.com/Samuel-Lewis/Password-Custard/app/models"
)

// SymbolSimple inserts a random symbol between words
func SymbolSimple(s []string) []string {
	return Insert(models.GetSymbol(), s)
}

// SmybolRandom inserts a random symbol into an exsiting word
func SmybolRandom(s []string) []string {
	// TODO
	return SymbolSimple(s)
}

// SymbolReplace replaces a random symbol in an exsiting word
func SymbolReplace(s []string) []string {
	// TODO
	return SymbolSimple(s)
}
