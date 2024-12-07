package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

func CommentHandeler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	temp, err := template.ParseFiles("/home/yjaouhar/Forum/internal/presentation/templates/comment/list.html")
	if err != nil {
		return
	}

	comment := r.FormValue("commant")
	count := 0
	for _, v := range comment {
		if v == '\n' {
			continue
		}
		count++
	}
	if comment == "" || count > 1000 {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	fmt.Println(comment)

	temp.Execute(w, nil)
}
