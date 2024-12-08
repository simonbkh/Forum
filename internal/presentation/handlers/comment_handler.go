package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"forum/internal/data/queries"
	"forum/internal/logic/services"
	"forum/internal/logic/utils"
	"forum/internal/presentation/templates"
)

func CommentHandeler(w http.ResponseWriter, r *http.Request) {
	id_post, _ := strconv.Atoi(r.PathValue("ID"))
	comment := r.FormValue("commant")
	date := time.Now().Format("2006-01-02 15:04:05")
	token, _ := r.Cookie("token")
	id_user, err := queries.GetIdUser(token.Value)
	if err != nil {
		return
	}
	er := queries.InsertComment(id_post, id_user, comment, date)
	if er != nil {
		return
	}
	utils.Check_Comment(comment)
	http.Redirect(w, r, fmt.Sprintf("/getcomment/%v", id_post), http.StatusSeeOther)
}

func GetCommment(w http.ResponseWriter, r *http.Request) {
	Comment, err := services.Comment_service(r.PathValue("ID"))
	if err != nil {
		return
	}

	templates.CommentTemplate.Execute(w, Comment)
}
