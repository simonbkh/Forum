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
	if !Presentation.IsLogged {
		Presentation.IsLogged = isTokenValid(r)
	}
	

	templates.HomeTemplate.Execute(w, Presentation)
	isLogged = false
}

func isTokenValid(r *http.Request) bool {
	tocken, err := r.Cookie("token")
	if err != nil || tocken.Value == "" {
		return false 
	}

	return queries.CheckToken_Prisent_or_not(tocken.Value)
}




	// if !Presentation.IsLogged {
	// 	tocken, err := r.Cookie("token")
	// 	if err != nil {
	// 		Presentation.IsLogged = false
	// 	} else {
	// 		if tocken.Value != "" {
	// 			//fmt.Println(tocken.Value)
	// 			if queries.CheckToken_Prisent_or_not(tocken.Value) {
	// 				Presentation.IsLogged = true
	// 			} else {
	// 				Presentation.IsLogged = false
	// 			}
	// 		}
	// 	}
	// }

//http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	