package handlers

import (
	"fmt"
	"net/http"

	"forum/internal/logic/services"
	"forum/internal/presentation/templates"
)

func Posts(w http.ResponseWriter, r *http.Request) {
	// http.Exict(w, r, "posts.html", nil)
	templates.PostTemplate.Execute(w, nil)
}

func HandlPost(w http.ResponseWriter, r *http.Request) {
	err := services.PostInfo(w, r)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
		return
	}
	isLogged = true
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
