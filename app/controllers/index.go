package controllers

import (
	"net/http"
)

// Index servers the home path index.html
func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "app/views/index.html")
}
