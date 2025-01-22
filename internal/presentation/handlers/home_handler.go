package handlers

import (
	"net/http"
	"time"

	"forum/internal/data/modles"
	"forum/internal/logic/services"
	"forum/internal/logic/utils"
	"forum/internal/presentation/templates"
)
type PageData struct {
	UserStatus bool
	Posts    []services.POST
}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "GET" {
	// 	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	// var Posts Post
	// }
	er := utils.CheckUserSession(r)
	if er != nil {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	if !modles.UserStatus {
		http.SetCookie(w, &http.Cookie{
			Name:    "SessionToken",
			Value:   "",
			Expires: time.Unix(0, 0),
		})
	}
	err := services.GetPosts(&services.Posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := PageData{
		UserStatus: modles.UserStatus,
		Posts:    services.Posts,
	}
	er = templates.HomeTemplate.Execute(w, data)
	if er != nil {
	}

}
