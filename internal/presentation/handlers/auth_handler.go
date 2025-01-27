package handlers

import (
	"fmt"
	"net/http"

	"forum/internal/logic/services"
	"forum/internal/logic/utils"
	"forum/internal/presentation/templates"
)

func Login(w http.ResponseWriter, r *http.Request) {
	templates.LoginTemplate.Execute(w, nil)
}

func Register(w http.ResponseWriter, r *http.Request) {
	templates.RegisterTemplate.Execute(w, nil)
}
func RegisterInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	err := services.Register_Service(w, r)
	if utils.IsErrors(err) {
		HandleError(w, err, http.StatusBadRequest)
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
		HandleError(w, err, http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "POST" {
	// 	http.Error(w, "metod not allowed", http.StatusMethodNotAllowed)
	// 	return
	// }
	er := services.Logout_Service(w, r)
	if utils.IsErrors(er) {
		http.Error(w, fmt.Sprintf("%v", er), http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
func HandleError(w http.ResponseWriter, err error, status int)  {
	type Error struct {
		ErrorCode    int
		ErrorMessage string
	}

	errorData := Error{
		ErrorCode:    status,
		ErrorMessage: err.Error(),
	}
	fmt.Println(errorData)
	templates.ErrorTemplate.Execute(w, errorData)
}