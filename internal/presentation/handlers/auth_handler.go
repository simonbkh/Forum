package handlers

import (
	"fmt"
	"net/http"
	"time"

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

func LoginInfo(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "POST" {
	// 	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	// }
	tocken, err := services.Login_Service(w, r)

	if utils.IsErrors(err) {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tocken,
		Expires: time.Now().Add(24 * time.Hour),
	})
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
