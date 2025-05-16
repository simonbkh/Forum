package middleware

import (
	"encoding/json"
	"net/http"
	"time"

	"forum/internal/data/modles"
	"forum/internal/data/queries"
)

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
