package handlers

import (
	"fmt"
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
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	filteredPosts , err := services.Category_Service(w, r)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
	}

	data := CatData {
		IsLogged: isLogged,
		Posts:    filteredPosts,
	}

	templates.HomeTemplate.Execute(w,data)

}
