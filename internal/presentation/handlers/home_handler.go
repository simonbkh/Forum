package handlers

import (
	"net/http"

	"forum/internal/presentation/templates"
)

var isLogged bool

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	templates.HomeTemplate.Execute(w, isLogged)
}
