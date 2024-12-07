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
	router.HandleFunc("/", handlers.HomeHandler)
	router.HandleFunc("/login", handlers.Login)
	router.HandleFunc("/static/css/{file}", handlers.Static)
	router.HandleFunc("/register", handlers.Register)
	router.HandleFunc("/loginInfo", handlers.LoginInfo)
	router.HandleFunc("/registerInfo", handlers.RegisterInfo)
	router.HandleFunc("/new_post", handlers.NewPostHandler)
	router.HandleFunc("/postinfo", handlers.PostInfo)
	router.HandleFunc("/logout", handlers.Logout)
	router.HandleFunc("/category/", handlers.CategoryHandler)
	router.HandleFunc("/comment", handlers.CommentHandeler)

	fmt.Println("website is running on: http://localhost:8080")

	err = http.ListenAndServe(":8080", router)
	if utils.IsErrors(err) {
		return err
	}
	return nil
}
