package handlers

import (
	"net/http"
	"time"

	models "forum/internal/data/database/modles"
	"forum/internal/data/queries"
	"forum/internal/presentation/templates"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "GET" {
	// 	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	// var Posts Post
	// }
	te, err := r.Cookie("SessionToken")
	if !models.UserStatus {
		if err != nil || te.Value == "" {
			models.UserStatus = false
		} else {
			bol, expiry := queries.IssesionidAvailable(te.Value, "")
			if bol && expiry.After(time.Now()) {
				models.UserStatus = true
			} else {
				err := queries.Removesesionid(te.Value, "")
				if err != nil {
					return
				}

				models.UserStatus = false
			}
		}
	}
	er := templates.HomeTemplate.Execute(w, models.UserStatus)
	if er != nil {
	}
	models.UserStatus = false
}
