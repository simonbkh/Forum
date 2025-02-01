package handlers

import (
	"net/http"

	"forum/internal/data/database"
	"forum/internal/data/modles"
	"forum/internal/logic/services"
	"forum/internal/presentation/templates"
)

var newPosts []database.Post

type PageData struct {
	UserStatus bool
	Posts      []services.POST
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// HandleError(w, nil, http.StatusNotFound)
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	err := services.GetPosts(&services.Posts) 
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	newPosts = services.TimeDifference(newPosts, services.Posts)
	data := PageData{
		UserStatus: modles.UserStatus,
		Posts:      newPosts,
	}
	er := templates.HomeTemplate.Execute(w, data)
	if er != nil {
		http.Error(w, er.Error(), http.StatusInternalServerError)
		return
	}
}
