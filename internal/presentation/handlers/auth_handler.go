package handlers

import (
	"fmt"
	"forum/internal/logic/services"
	"forum/internal/logic/utils"
	"forum/internal/presentation/templates"
	"net/http"
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
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	err := services.Login_Service(w, r)
	if utils.IsErrors(err) {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
		return
	}
	Getcoockes(r)
	// cookie := &http.Cookie{
	// 	Name:  "email",
	// 	Value: "12@",
	// }
	// http.SetCookie(w, cookie)
	er := templates.HomeTemplate.Execute(w, nil)
	if er != nil {
		return
	}
	// http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func Getcoockes(r *http.Request) {
	emailCookie, err := r.Cookie("email")
	if err != nil {
		fmt.Println("Error reading email cookie:", err)
		return
	}
	fmt.Println(emailCookie)
}
