package handlers

import (
	"fmt"
	"net/http"

	"forum/internal/presentation/templates"
)

func Post(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }
	templates.PostTemplate.Execute(w, nil)
}

func PostInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }
	content := r.FormValue("post")
	categories := r.Form["category"]
	if len(categories) == 0 {
		categories = append(categories, "general")
	}
	fmt.Println(content, categories)

	http.Redirect(w,r,"/",http.StatusSeeOther)
}
