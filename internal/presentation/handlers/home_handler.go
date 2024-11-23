package handlers

import (
	"net/http"

	"net/http"

	"forum/internal/presentation/templates"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	templates.HomeTemplate.Execute(w,nil)
}
