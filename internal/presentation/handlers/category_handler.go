package handlers

import (
	"fmt"
	"net/http"

	"forum/internal/data/modles"
	"forum/internal/logic/services"
	"forum/internal/presentation/templates"
)

type CatData struct {
	UserStatus bool
	Posts    []services.POST
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

	data := CatData{
		UserStatus: modles.UserStatus,
		Posts:    filteredPosts,
	}

	templates.HomeTemplate.Execute(w, data)
}
