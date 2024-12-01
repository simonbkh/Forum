package handlers

import (
	"fmt"
	"net/http"

	"forum/internal/presentation/templates"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "GET" {
	// 	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	// var Posts Post
	// }
	Text := "tt"
	// te, err := r.Cookie("SessionToken")

	// if err != nil {
	// 	Text = "tt"
	// } else {
	// 	fmt.Println(te.Value)
	// 	if queries.IssesionidAvailable(te.Value) {
	// 		Text = "tt"
	// 	} else {
	// 		Text = "ddddddd"
	// 	}
	// }
	er := templates.HomeTemplate.Execute(w, Text)
	if er != nil {
		fmt.Println(er)
	}
}
