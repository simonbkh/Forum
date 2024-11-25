package handlers

import (
	"net/http"

	"forum/internal/logic/services"
	"forum/internal/presentation/templates"
)
type PageData struct {
	IsLogged bool
	Posts []services.Post
}
var isLogged bool

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		IsLogged: isLogged,
		Posts: services.Posts,
	}
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	templates.HomeTemplate.Execute(w, data)
}
