package handlers

import (
	"fmt"
	"net/http"

	"forum/internal/presentation/templates"
)

func Post(w http.ResponseWriter, r *http.Request) {
	templates.PostTemplate.Execute(w, nil)
}

func PostInfo(w http.ResponseWriter, r *http.Request) {
	content := r.FormValue("post")
	categories := r.Form["category"]
	if len(categories) == 0 {
		categories = append(categories, "general")
	}
	fmt.Println(content, categories)
}
