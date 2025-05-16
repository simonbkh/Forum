package handlers

import (
	"net/http"

	"forum/internal/data/database"
	"forum/internal/data/queries"
	"forum/internal/logic/services"
	"forum/internal/presentation/templates"
)

type Mypost struct {
	UserStatus bool
	Posts      []database.Post
}
type Create struct {
	Err error
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/post" {
		HandleError(w, 404)
		return
	}
	if r.Method != "GET" {
		HandleError(w, http.StatusMethodNotAllowed)
		return
	}

	templates.Create_post.Execute(w, nil)
}

func SubmittedPost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/submit-post" {
		HandleError(w, 404)
		return
	}
	if r.Method != "POST" {
		HandleError(w, 405)
		return
	}
	// fmt.Println(r.Body)
	err := services.Post_Service(w, r)
	if err != nil {
		data := Create{
			Err: err,
		}
		templates.Create_post.Execute(w, data)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func MyPosts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/myPosts" {
		HandleError(w, 404)
		return
	}
	token, err := r.Cookie("SessionToken")
	if err != nil {
		HandleError(w, http.StatusBadRequest)
	}

	user_id, err := queries.GetId(token.Value)
	if err != nil {
		HandleError(w, http.StatusBadRequest)
	}
	NewPosts, err := queries.GetPost(user_id)
	if err != nil {
		HandleError(w, http.StatusBadRequest)
	}

	templates.HomeTemplate.Execute(w, Page(r, NewPosts))
}

func LikedPosts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/likedPosts" {
		HandleError(w, 404)
		return
	}
	token, err := r.Cookie("SessionToken")
	if err != nil {
		HandleError(w, http.StatusBadRequest)
	}

	user_id, err := queries.GetId(token.Value)
	if err != nil {
		HandleError(w, http.StatusBadRequest)
	}

	LikedPosts, err := queries.GetLikedPosts(user_id)
	if err != nil {
		HandleError(w, http.StatusBadRequest)
	}

	templates.HomeTemplate.Execute(w, Page(r, LikedPosts))
	// fmt.Println(err)
}
