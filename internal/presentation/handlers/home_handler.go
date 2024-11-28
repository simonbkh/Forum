package handlers

import (
	"fmt"
	"net/http"

	"forum/internal/data/queries"
	"forum/internal/presentation/templates"
)

type PageData struct {
	IsLogged bool
	Posts    []queries.Post
}


var isLogged bool

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	Presentation := PageData{
		IsLogged: isLogged,
	}
	var err error
	Presentation.Posts, err = queries.GetPosts()
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
		return
	}
	// if r.Method != "GET" {
	// 	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	// }

	templates.HomeTemplate.Execute(w, Presentation)
	isLogged = false
}
