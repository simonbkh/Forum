package handlers

import (
	"fmt"
	"net/http"

	"forum/internal/logic/services"
	"forum/internal/logic/validators"
	"forum/internal/presentation/templates"
)

func Posts(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		templates.ErrorTemlate.Execute(w, "Method Not Allowed..")
		return
	}
	if !validators.Check_cokes(r) {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	// http.Exict(w, r, "posts.html", nil)
	templates.PostTemplate.Execute(w, nil)
}

func HandlPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		templates.ErrorTemlate.Execute(w, "Method Not Allowed..")
		//http.Error(w, "Method Not Allowed.", http.StatusMethodNotAllowed)
		return
	}

	err := services.PostInfo(w, r)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
