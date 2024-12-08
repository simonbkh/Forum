package handlers

import (
	"net/http"

	"forum/internal/logic/services"
	"forum/internal/logic/validators"
	"forum/internal/presentation/templates"
)

func NewPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		Errore(w, http.StatusMethodNotAllowed)
	}
	_, err := validators.Allowed(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}
	templates.PostTemplate.Execute(w, nil)
}

func PostInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		Errore(w, http.StatusMethodNotAllowed)

	}
	err := services.Post_Service(w, r)
	if err != nil {
		Errore(w, http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
