package utils

import (
	"fmt"
	"net/http"
	"time"
)

func SetTokenCookie(w http.ResponseWriter, token string) {
	// Create a new cookie
	fmt.Println(token)
	var cookie *http.Cookie
	if token == "" {
		cookie = &http.Cookie{
			Name:    "token",
			Value:   token,
			Expires: time.Now().Add(-1),
			Path:    "/", // Cookie is valid for all paths
		}
	} else {
		cookie = &http.Cookie{
			Name:    "token",
			Value:   token,
			Expires: time.Now().Add(24 * time.Hour), // Expires in 24 hours
			Path:    "/",                            // Cookie is valid for all paths
		}
	}
	http.SetCookie(w, cookie)
}
