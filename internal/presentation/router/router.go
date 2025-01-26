package router

import (
	"fmt"
	"net/http"
	"time"

	"forum/internal/data/modles"
	"forum/internal/data/queries"
	"forum/internal/logic/utils"
	"forum/internal/presentation/handlers"
	"forum/internal/presentation/templates"
)

func Router(router *http.ServeMux) error {
	var err error
	err = templates.ParseFiles()
	if utils.IsErrors(err) {
		fmt.Println(err)
		return err
	}
	// hhh:= Middleware(http.HandlerFunc(handlers.HomeHandler))
	router.Handle("/", Middleware(http.HandlerFunc(handlers.HomeHandler)))
	router.Handle("/login", Middleware(http.HandlerFunc(handlers.Login)))
	router.Handle("/static/css/{file}", http.HandlerFunc(handlers.Static))
	router.Handle("/static/js/{file}", http.HandlerFunc(handlers.JS))
	router.Handle("/forum/internal/presentation/static/js/{file}", http.HandlerFunc(handlers.JS))
	router.Handle("/static/images/{file}", http.HandlerFunc(handlers.Image))
	router.Handle("/loginInfo", http.HandlerFunc(handlers.LoginInfo))
	router.Handle("/registerInfo", http.HandlerFunc(handlers.RegisterInfo))
	router.Handle("/logout", Middleware(http.HandlerFunc(handlers.Logout)))
	router.Handle("/post", Middleware(http.HandlerFunc(handlers.PostHandler)))
	router.Handle("/submit-post", Middleware(http.HandlerFunc(handlers.SubmittedPost)))
	router.Handle("/category/", Middleware(http.HandlerFunc(handlers.CategoryHandler)))
	router.Handle("/myPosts", Middleware(http.HandlerFunc(handlers.MyPosts)))
	router.HandleFunc("/newcomment", handlers.Creatcomment)
	router.HandleFunc("/getcomment", handlers.GetComment)
	fmt.Println("website is running on: http://localhost:8081")
	err = http.ListenAndServe(":8081", router)
	if utils.IsErrors(err) {
		return err
	}
	return nil
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		te, err := r.Cookie("SessionToken")
		if err != nil || te.Value == "" {
			modles.UserStatus = false
			if r.URL.Path != "/" && r.URL.Path != "/login" && r.URL.Path != "/register" && r.URL.Path != "/static/js/{file}" {
				http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
				return

			}
			next.ServeHTTP(w, r)
			return
		}

		bol, expiry := queries.IssesionidAvailable(te.Value, "")
		if !bol || expiry.Before(time.Now()) {
			modles.UserStatus = false
			if r.URL.Path != "/"  {
				http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
				return
			}
			next.ServeHTTP(w, r)
			return
		}
		modles.UserStatus = true
		next.ServeHTTP(w, r)
	})
}
