package router

import (
	"fmt"
	"net/http"

	"forum/internal/logic/utils"
	"forum/internal/presentation/handlers"
	"forum/internal/presentation/templates"
)

func Router(router *http.ServeMux) error {
	var err error
	err = templates.ParseFiles()
	if utils.IsErrors(err) {
		return err
	}
	router.HandleFunc("GET /", handlers.HomeHandler)
	router.HandleFunc("GET /login", handlers.Login)
	router.HandleFunc("GET /post", handlers.Post)
	router.HandleFunc("GET /static/css/{file}", handlers.Static)
	router.HandleFunc("GET /register", handlers.Register)
	router.HandleFunc("POST /loginInfo", handlers.LoginInfo)
	router.HandleFunc("POST /registerInfo", handlers.RegisterInfo)

	fmt.Println("website is running on: http://localhost:8080")

	err = http.ListenAndServe(":8080", router)
	if utils.IsErrors(err) {
		return err
	}
	return nil
}
