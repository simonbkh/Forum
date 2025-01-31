package handlers

import (
	"encoding/json"
	"net/http"

	"forum/internal/data/queries"
	"forum/internal/logic/utils"
)

type newcom struct {
	Post_id string `json:"post"`
	Comment string `json:"comment"`
	Date    string `json:"date"`
}

func Creatcomment(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		sesiontoken, _ := r.Cookie("SessionToken")
		User_id, err := queries.GetId(sesiontoken.Value)
		if err != nil {
			HandleError(w, err, http.StatusUnauthorized)
		}
		var info newcom
		errore := json.NewDecoder(r.Body).Decode(&info)
		if errore != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		if info.Comment != "" && len(info.Comment) <= 1000 {
			er := queries.InsertComment(User_id, utils.Convstr(info.Post_id), info.Comment, info.Date)
			if er != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		}

	}
}

type get struct {
	Post_id string `json:"post"`
}

func GetComment(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		var comm get

		err := json.NewDecoder(r.Body).Decode(&comm)
		if err != nil {

			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		comment, er := queries.GetCommment(utils.Convstr(comm.Post_id))

		if er != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(comment); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

	}
}
func GetLenComment(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		var comm get

		err := json.NewDecoder(r.Body).Decode(&comm)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		comment, er := queries.GetCommment(utils.Convstr(comm.Post_id))
		if er != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(len(comment)); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

	}
}
