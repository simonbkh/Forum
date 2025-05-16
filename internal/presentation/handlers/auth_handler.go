package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"forum/internal/logic/services"
	"forum/internal/presentation/templates"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		HandleError(w, 405)
		return
	}
	if r.URL.Path != "/login" {
		HandleError(w, 404)
		return
	}
	err := templates.LoginTemplate.Execute(w, nil)
	if err != nil {
		HandleError(w, 500)
		return
	}
}

func RegisterInfo(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/registerInfo" {
		HandleError(w, 404)
		return
	}

	if r.Method != "POST" {
		HandleError(w, 405)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := services.Register_Service(w, r)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "message": err.Error()})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "message": ""})
}

func LoginInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		HandleError(w, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/loginInfo" {
		HandleError(w, 404)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err := services.Login_Service(w, r)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "message": err.Error()}) // json.NewEncoder(w).Encode(err.type)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "message": ""})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	er := services.Logout_Service(w, r)
	if er != nil {
		http.Error(w, fmt.Sprintf("%v", er), http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func HandleError(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	type Error struct {
		ErrorCode    int
		ErrorMessage string
	}
	errorData := Error{
		ErrorCode:    status,
		ErrorMessage: http.StatusText(status),
	}
	err := templates.ErrorTemplate.Execute(w, errorData)
	if err != nil {
		http.Error(w, "Internal server error", 500)
		return
	}
}
