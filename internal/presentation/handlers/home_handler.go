package handlers

import (
	"net/http"
	"time"

	"forum/internal/data/queries"
	"forum/internal/logic/services"
	"forum/internal/presentation/templates"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "GET" {
	// 	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	// var Posts Post
	// }

	if services.Text == "tt" {
		te, err := r.Cookie("SessionToken")
		if err != nil || te.Value == "" {
			services.Text = "tt"
		} else {
			bol, expiry := queries.IssesionidAvailable(te.Value, "")
			if bol && expiry.After(time.Now()) {
				services.Text = ""
			} else {
				err := queries.Removesesionid(te.Value, "")
				if err != nil {
					return
				}
				services.Text = "tt"
			}
		}
	}
	er := templates.HomeTemplate.Execute(w, services.Text)
	if er != nil {
	}
	services.Text = "tt"
}
