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
	er := templates.HomeTemplate.Execute(w,Text)
	if er != nil {
		fmt.Println(er)
	}
}
