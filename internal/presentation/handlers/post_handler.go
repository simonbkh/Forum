package handlers

import (
	"net/http"

	"forum/internal/data/database"
	"forum/internal/data/queries"
	"forum/internal/logic/services"
	"forum/internal/presentation/templates"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// _, err := validators.Allowed(w, r)
	// if err != nil {
	// 	// http.Error(w, "u cant lol!!", http.StatusUnauthorized)
	// 	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
	// 	retur
	// }
	templates.Create_post.Execute(w, nil)
}

func SubmittedPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	err := services.Post_Service(w, r)
	if err != nil {
		HandleError(w, err, http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

var Mypost struct {
	Posts []database.Post
}

func MyPosts(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("SessionToken")
	if err != nil {
		HandleError(w, err, http.StatusBadRequest)
	}

	user_id, err := queries.GetId(token.Value)
	if err != nil {
		HandleError(w, err, http.StatusBadRequest)
	}
	// NewPosts := []database.Post{}
	NewPosts, err := queries.GetPost(user_id)
	if err != nil {
		HandleError(w, err, http.StatusBadRequest)
	}

	Mypost.Posts = NewPosts
	templates.MyPossts.Execute(w, Mypost)

	// return NewPosts
}
