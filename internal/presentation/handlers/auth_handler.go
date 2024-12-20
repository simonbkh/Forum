package handlers

import (
	"fmt"
	"net/http"

	"forum/internal/logic/services"
	"forum/internal/logic/utils"
	"forum/internal/presentation/templates"
)

//      /login, /register routes


func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	templates.RegisterTemplate.Execute(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	templates.LoginTemplate.Execute(w, nil)
}

func RegisterInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	err := services.Register_Service(w, r)
	if utils.IsErrors(err) {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func LoginInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	err := services.Login_Service(w, r)
	if utils.IsErrors(err) {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
		return
	}
	isLogged = true
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	err := services.Logout(w  ,r )
	if utils.IsErrors(err) {
		//maeereftch wach hadak howa status code
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
	}
	isLogged = false
	http.Redirect(w, r, "/", http.StatusSeeOther)
}