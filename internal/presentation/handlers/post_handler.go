package handlers

import (
	"fmt"
	"net/http"

	"forum/internal/presentation/templates"
)

func Post(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "POST" {
	// 	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	// }
	er := templates.Create_post.Execute(w, nil)
	if er != nil {
		fmt.Println(er)
		return
	}
}
