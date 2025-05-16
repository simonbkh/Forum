package router

import (
	"fmt"
	"net/http"

	"forum/internal/presentation/handlers"
	"forum/internal/presentation/middleware"
	"forum/internal/presentation/templates"
)

func Router(router *http.ServeMux) error {
	var err error
	err = templates.ParseFiles()
	if err!=nil{
		return err
	}
	router.Handle("/", middleware.Middleware(http.HandlerFunc(handlers.HomeHandler)))
	router.Handle("/api/{direction}/{id}/reaction", middleware.Middleware(http.HandlerFunc(handlers.HandleReaction)))
	router.Handle("/likedPosts", middleware.Middleware(http.HandlerFunc(handlers.LikedPosts)))
	router.Handle("/login", middleware.Middleware(http.HandlerFunc(handlers.Login)))
	router.Handle("/static/{folder}/{file}", http.HandlerFunc(handlers.Static))
	router.Handle("/loginInfo", http.HandlerFunc(handlers.LoginInfo))
	router.Handle("/registerInfo", http.HandlerFunc(handlers.RegisterInfo))
	router.Handle("/logout", middleware.Middleware(http.HandlerFunc(handlers.Logout)))
	router.Handle("/post", middleware.Middleware(http.HandlerFunc(handlers.PostHandler)))
	router.Handle("/submit-post", http.HandlerFunc(handlers.SubmittedPost))
	router.Handle("/category/", (http.HandlerFunc(handlers.CategoryHandler)))
	router.Handle("/myPosts", middleware.Middleware(http.HandlerFunc(handlers.MyPosts)))
	router.Handle("/newcomment", middleware.Middleware(http.HandlerFunc(handlers.Creatcomment)))
	router.HandleFunc("/getcomment", handlers.GetComment)
	router.HandleFunc("/commentlen", handlers.GetLenComment)

	fmt.Println("website is running on: http://localhost:8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		return err
	}
	return nil
}
