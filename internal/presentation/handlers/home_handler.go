package handlers

import (
	"fmt"
	"net/http"

	"forum/internal/presentation/templates"
)

type Post struct {
	Text string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "GET" {
	// 	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	var Posts Post
	// }
	Posts.Text = "tt"
	er := templates.HomeTemplate.Execute(w, Posts)
	if er != nil {
		fmt.Println(er)
	}
}
