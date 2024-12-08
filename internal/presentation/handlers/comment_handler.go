package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"forum/internal/data/queries"
	"forum/internal/logic/services"
)

func CommentHandeler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	id_post, _ := strconv.Atoi(r.PathValue("ID"))
	comment := r.FormValue("commant")
	date := time.Now().Format("2006-01-02 15:04:05")
	token, _ := r.Cookie("token")
	id_user, err := queries.GetIdUser(token.Value)
	if err != nil {
		return
	}

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
	er := queries.InsertComment(id_post, id_user, comment, date)
	if er != nil {
		fmt.Println(er)
		return
	}
}

func GetCommment(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("ID"))

	if err != nil || id < 0 || id > len(services.Posts) {
		return
	}
	temp, err := template.ParseFiles("../internal/presentation/templates/comment/list.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	er := temp.Execute(w, services.Posts[len(services.Posts)-id])
	if er != nil {
		return
	}
}
