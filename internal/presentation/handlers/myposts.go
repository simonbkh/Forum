package handlers

import (
	"net/http"

	"forum/internal/logic/services"
	"forum/internal/logic/validators"
	"forum/internal/presentation/templates"
)

func Myposts(w http.ResponseWriter, r *http.Request) {
	user_id, err := validators.Allowed(w, r)
	if err != nil {
		// redirect or smtg
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	posts := services.UserPosts(user_id)
	data := CatData{
		IsLogged: isLogged,
		Posts:    posts,
	}

	templates.HomeTemplate.Execute(w, data)
	// fmt.Println(posts)
}
