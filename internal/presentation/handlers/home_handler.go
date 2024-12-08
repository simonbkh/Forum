package handlers

import (
	"net/http"

	"forum/internal/logic/services"
	"forum/internal/presentation/templates"
)

type PageData struct {
	IsLogged bool
	Posts    []services.POST
}

var isLogged bool

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		Errore(w, http.StatusMethodNotAllowed)

	}
	err := services.GetPosts(&services.Posts)
	if err != nil {
		Errore(w, http.StatusInternalServerError)

	}
	data := PageData{
		IsLogged: isLogged,
		Posts:    services.Posts,
	}
	templates.HomeTemplate.Execute(w, data)
}
