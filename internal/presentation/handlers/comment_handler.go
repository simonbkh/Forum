package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"forum/internal/data/queries"
	utils "forum/internal/logic/Utils"
)

type newcom struct {
	Post_id string `json:"post"`
	User_id string `json:"id"`
	Comment string `json:"comment"`
	Date    string `json:"date"`
}

func Creatcomment(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		var info newcom

		err := json.NewDecoder(r.Body).Decode(&info)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		fmt.Println("------", info)
		er := queries.InsertComment(utils.Convstr(info.User_id), utils.Convstr(info.Post_id), info.Comment, info.Date)
		if er != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
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
		//
		if err := json.NewEncoder(w).Encode(comment); err != nil {

			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

	}
}
