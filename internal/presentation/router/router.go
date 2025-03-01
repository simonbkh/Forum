package router

import (
	"encoding/json"
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
	router.Handle("/api/posts/{post_id}/reaction", Middleware(http.HandlerFunc(handlers.HandleReaction)))
	router.Handle("/login", Middleware(http.HandlerFunc(handlers.Login)))
	router.Handle("/static/css/{file}", http.HandlerFunc(handlers.Static))
	router.Handle("/static/js/{file}", http.HandlerFunc(handlers.JS))
	router.Handle("/static/images/{file}", http.HandlerFunc(handlers.Image))
	router.Handle("/loginInfo", http.HandlerFunc(handlers.LoginInfo))
	router.Handle("/registerInfo", http.HandlerFunc(handlers.RegisterInfo))
	router.Handle("/logout", Middleware(http.HandlerFunc(handlers.Logout)))
	router.Handle("/post", Middleware(http.HandlerFunc(handlers.PostHandler)))
	router.Handle("/submit-post", Middleware(http.HandlerFunc(handlers.SubmittedPost)))
	router.Handle("/category/", (http.HandlerFunc(handlers.CategoryHandler)))
	router.Handle("/myPosts", Middleware(http.HandlerFunc(handlers.MyPosts)))
	router.Handle("/newcomment", Middleware(http.HandlerFunc(handlers.Creatcomment)))
	router.HandleFunc("/getcomment", handlers.GetComment)
	router.HandleFunc("/commentlen", handlers.GetLenComment)
	fmt.Println("website is running on: http://localhost:8080")
	err = http.ListenAndServe(":8080", router)
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
			if CheckUserSession(r, w) {
				next.ServeHTTP(w, r)
				return
			}
			return
		}

		bol, expiry := queries.IssesionidAvailable(te.Value, "")
		if !bol || expiry.Before(time.Now()) {
			modles.UserStatus = false
			if CheckUserSession(r, w) {
				next.ServeHTTP(w, r)
				return
			}
			return
		}
		fmt.Println("user is logged in",r.URL.Path)
		modles.UserStatus = true
		next.ServeHTTP(w, r)
	})
}
 
func CheckUserSession(r *http.Request, w http.ResponseWriter) bool {
	if r.Header.Get("Content-Type") == "application/json" {
		if r.URL.Path != "/" && r.URL.Path != "/login" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "message": "unauthorized"})
			return false
		}
		return true
	}
	if r.URL.Path != "/" && r.URL.Path != "/login" {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return false
	}
	return true
}
