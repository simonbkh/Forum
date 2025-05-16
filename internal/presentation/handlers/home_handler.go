package handlers

import (
	"net/http"
	"strconv"

	"forum/internal/data/database"
	"forum/internal/data/modles"
	"forum/internal/logic/services"
	"forum/internal/presentation/templates"
)

var newPosts []database.Post

type PageData struct {
	Path        string
	UserStatus  bool
	Posts       []services.POST
	CurrentPage int
	TotalPages  int
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		HandleError(w, http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		HandleError(w, http.StatusMethodNotAllowed)
		return
	}

	token := ""
	if len(r.Header.Values("cookie")) != 0 {
		token = r.Header.Values("cookie")[0][13:]
	}
	err := services.GetPosts(&services.Posts, token)
	if err != nil {
		HandleError(w, http.StatusInternalServerError)
		return
	}
	err = templates.HomeTemplate.Execute(w, Page(r, []database.Post{}))
	if err != nil {
		HandleError(w, 400)
		return
	}
}

func Page(r *http.Request, post []database.Post) PageData {
	totalPosts := 0
	pageStr := r.URL.Query().Get("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	if len(post) != 0 {
		totalPosts = len(post)
	} else {
		totalPosts = len(services.Posts)
	}

	totalPages := (totalPosts + 10 - 1) / 10

	if page > totalPages {
		page = totalPages
	}
	start := (page - 1) * 10
	end := start + 10
	if end > totalPosts {
		end = totalPosts
	}
	var paginatedPosts []database.Post
	if len(post) != 0 {
		paginatedPosts = post[start:end]
	} else {
		paginatedPosts = services.Posts[start:end]
	}

	newPosts = services.TimeDifference(paginatedPosts)
	data := PageData{
		UserStatus:  modles.UserStatus,
		Posts:       newPosts,
		CurrentPage: page,
		TotalPages:  totalPages,
	}
	return data
}
