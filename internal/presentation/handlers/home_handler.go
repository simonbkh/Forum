package handlers

import (
	"net/http"

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
			if queries.IssesionidAvailable(te.Value) {
				services.Text = ""
			} else {
				services.Text = "tt"
			}
		}
	}
	er := templates.HomeTemplate.Execute(w, services.Text)
	if er != nil {
	}
	services.Text = "tt"
}
