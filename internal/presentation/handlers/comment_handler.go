package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"forum/internal/data/queries"
)

type newcom struct {
	Post_id string `json:"post"`
	Comment string `json:"comment"`
	Date    string `json:"date"`
}
type get struct {
	Post_id string `json:"post"`
}

func Creatcomment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		HandleError(w, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/newcomment" {
		HandleError(w, http.StatusNotFound)
		return
	}
	sesiontoken, _ := r.Cookie("SessionToken")
	User_id, err := queries.GetId(sesiontoken.Value)
	if err != nil {
		HandleError(w, 500)
		return
	}
	var info newcom
	errore := json.NewDecoder(r.Body).Decode(&info)
	if errore != nil {
		HandleError(w, 500)
		return
	}

	if info.Comment != "" && len(info.Comment) <= 2000 && strings.TrimSpace(info.Comment) != "" {
		postid, err := strconv.Atoi(info.Post_id)
		if err != nil || !queries.ValidPostId(postid) {
			HandleError(w, 400)
			return
		}
		er := queries.InsertComment(User_id, postid, info.Comment, info.Date)
		if er != nil {
			HandleError(w, 500)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "message": ""})

	} else {
		HandleError(w, 404)
		return
	}
}

func GetComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		HandleError(w, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/getcomment" {
		HandleError(w, 404)
		return
	}
	token := ""
	var comm get
	err := json.NewDecoder(r.Body).Decode(&comm)
	if err != nil {
		HandleError(w, 500)
		return
	}
	postid, err := strconv.Atoi(comm.Post_id)
	if err != nil {
		HandleError(w, 400)
		return
	}
	if !queries.ValidPostId(postid) {
		HandleError(w, 400)
		return
	}
	if len(r.Header.Values("cookie")) != 0 {
		token = r.Header.Values("cookie")[0][13:] // might wanna protect the check
	}
	comment, er := queries.GetCommment(postid, token)
	if er != nil {
		HandleError(w, 500)
		return
	}
	if err := json.NewEncoder(w).Encode(comment); err != nil {
		HandleError(w, 500)
		return
	}
}

func GetLenComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		HandleError(w, 405)
		return
	}
	if r.URL.Path != "/commentlen" {
		HandleError(w, 404)
		return
	}
	var comm get
	err := json.NewDecoder(r.Body).Decode(&comm)
	if err != nil {
		HandleError(w, 500)
		return
	}
	postid, err := strconv.Atoi(comm.Post_id)
	if err != nil {
		HandleError(w, 400)
		return
	}
	if !queries.ValidPostId(postid) {
		HandleError(w, 400)
		return
	}
	comment, er := queries.GetCommment(postid, "")
	if er != nil {
		HandleError(w, 500)
		return
	}
	if err := json.NewEncoder(w).Encode(len(comment)); err != nil {
		HandleError(w, 500)
		return
	}
}
