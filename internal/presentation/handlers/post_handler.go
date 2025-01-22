package handlers

import (
	"net/http"

	"forum/internal/logic/services"
	"forum/internal/presentation/templates"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// _, err := validators.Allowed(w, r)
	// if err != nil {
	// 	// http.Error(w, "u cant lol!!", http.StatusUnauthorized)
	// 	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
	// 	retur
	// }
	templates.Create_post.Execute(w, nil)
}

func SubmittedPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	err := services.Post_Service(w, r)
	if err != nil {
		HandleError(w, err, http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
