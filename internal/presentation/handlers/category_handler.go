package handlers

import (
	"forum/internal/logic/services"
	"forum/internal/presentation/templates"
	"net/http"
)

type CatData struct {
	IsLogged bool
	Posts    []services.POST
}

func CategoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		Errore(w, http.StatusMethodNotAllowed)

	}
	filteredPosts, err := services.Category_Service(w, r)
	if err != nil {
		Errore(w, http.StatusBadRequest)
	}

	data := CatData{
		IsLogged: isLogged,
		Posts:    filteredPosts,
	}

	templates.HomeTemplate.Execute(w, data)

}
