package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"forum/internal/logic/services"
	"forum/internal/logic/utils"
	"forum/internal/presentation/templates"
)

func Login(w http.ResponseWriter, r *http.Request) {
	templates.LoginTemplate.Execute(w, nil)
}

// func Register(w http.ResponseWriter, r *http.Request) {
// 	templates.RegisterTemplate.Execute(w, nil)
// }
func RegisterInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	w.Header().Set("Content-Type", "application/json")
	err := services.Register_Service(w, r)
	if utils.IsErrors(err) {
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false,"message": err.Error(),})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true,"message": ""})
}

func LoginInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

	w.Header().Set("Content-Type", "application/json")	

	err := services.Login_Service(w, r)
	if utils.IsErrors(err) {
	json.NewEncoder(w).Encode(map[string]interface{}{"success": false,"message": err.Error(),}) //json.NewEncoder(w).Encode(err.type)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true,"message": ""})

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