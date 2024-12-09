package handlers

import (
	"net/http"
	"time"

	"forum/internal/data/modles"
	"forum/internal/logic/utils"
	"forum/internal/presentation/templates"
)

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
	er = templates.HomeTemplate.Execute(w, modles.UserStatus)
	if er != nil {
	}

}
