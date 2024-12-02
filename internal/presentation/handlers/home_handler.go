package handlers

import (
	"fmt"
	"net/http"

	"forum/internal/data/queries"
	"forum/internal/logic/validators"
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
	if !Presentation.IsLogged {
		Presentation.IsLogged = validators.IsTokenValid(r)
	}

	templates.HomeTemplate.Execute(w, Presentation)
	
	isLogged = false
}
