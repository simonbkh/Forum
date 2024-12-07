package handlers

import (
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
	if !Presentation.IsLogged {
		Presentation.IsLogged = validators.IsTokenValid(r)
	}
	var err error

	Presentation.Posts, err = queries.GetPosts()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		templates.ErrorTemlate.Execute(w, err)
		//Presentation.IsLogged = false
		return
	}

	templates.HomeTemplate.Execute(w, Presentation) 

	isLogged = false
}
