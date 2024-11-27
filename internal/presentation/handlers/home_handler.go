package handlers

import (
	"net/http"

	"forum/internal/presentation/templates"
)

type PageData struct {
	IsLogged bool
	// Posts    []services.POST
}

var isLogged bool

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	Presentation := PageData{
		IsLogged: isLogged,
	}
	// if r.Method != "GET" {
	// 	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	// }

	templates.HomeTemplate.Execute(w, Presentation)
	isLogged = false
}
