package handlers

import (
	"fmt"
	"net/http"

	"forum/internal/logic/services"
	"forum/internal/logic/utils"
	"forum/internal/presentation/templates"
)

// /login, /register routes
func Register(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "GET" {
	// 	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	// }
	templates.RegisterTemplate.Execute(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "GET" {
	// 	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	// }
	templates.LoginTemplate.Execute(w, nil)
}

func RegisterInfo(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "POST" {
	// 	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	// }
	err := services.Register_Service(w, r)
	if utils.IsErrors(err) {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
}

type Users struct {
	Username string
	Text     string
}

func LoginInfo(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "POST" {
	// 	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	// }
	err := services.Login_Service(w, r)
	if utils.IsErrors(err) {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
		return
	}
	var User Users
	er := templates.HomeTemplate.Execute(w, User)
	if er != nil {
		fmt.Println("ooooooooook")
		return
	}
	// http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
