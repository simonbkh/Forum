package handlers

import (
	"fmt"
	"net/http"

	"forum/internal/logic/services"
	"forum/internal/presentation/templates"
)

func CategoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		HandleError(w, 405)
		return
	}
	filteredPosts, err := services.Category_Service(w, r)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
	}
	templates.HomeTemplate.Execute(w, Page(r, filteredPosts))
}
