package handlers

import (
	"fmt"
	"net/http"

	"forum/internal/data/database"
	"forum/internal/data/modles"
	"forum/internal/logic/services"
	"forum/internal/presentation/templates"
)

var catPosts []database.Post

type CatData struct {
	UserStatus bool
	Posts      []services.POST
}

func CategoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	// h := strings.Split(r.URL.Path, "/")
	// fmt.Println(h)
	// url := strings.Join(h[2:], "/")
	// r.URL.Path = url
	filteredPosts, err := services.Category_Service(w, r)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
	}
	catPosts = services.TimeDifference(catPosts,filteredPosts)
	data := CatData{
		UserStatus: modles.UserStatus,
		Posts:      catPosts,
	}

	templates.HomeTemplate.Execute(w, data)
}
