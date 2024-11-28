package handlers

import (
	"fmt"
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
	 err := services.GetPosts(&services.Posts)
	 if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	 }
	data := PageData{
		IsLogged: isLogged,
		Posts:    services.Posts,
	}
	fmt.Println(services.Posts,isLogged)
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	templates.HomeTemplate.Execute(w, data)
}
