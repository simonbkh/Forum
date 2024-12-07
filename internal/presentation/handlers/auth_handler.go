package handlers

import (
	"net/http"
	"time"

	"forum/internal/logic/services"
	"forum/internal/presentation/templates"
)

// /login, /register routes
func Register(w http.ResponseWriter, r *http.Request) {
	if !services.Method(w, r, "GET") {
		return
	}
	templates.RegisterTemplate.Execute(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if !services.Method(w, r, "GET") {
		return
	}
	templates.LoginTemplate.Execute(w, nil)
}

func RegisterInfo(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.Method)
	if !services.Method(w,r,"POST"){
		return
	}
	err := services.Register_Service(w, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		templates.ErrorTemlate.Execute(w, err)
		return
	}
	isLogged = false
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func LoginInfo(w http.ResponseWriter, r *http.Request) {
	if !services.Method(w, r, "POST") {
		return
	}

	tocken, err := services.Login_Service(w, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		templates.ErrorTemlate.Execute(w, err)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tocken,
		Expires: time.Now().Add(24 * time.Hour),
	})
	isLogged = true
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func Log_out(w http.ResponseWriter, r *http.Request) {
	err := services.Log_out_Service(w, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		templates.ErrorTemlate.Execute(w, err)
		return
	}
	isLogged = false
	http.SetCookie(w, &http.Cookie{
		Name:   "token",
		Value:  "",
		MaxAge: -1,
	})
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
