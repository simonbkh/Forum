package utils

import (
	"net/http"
	"time"
)

func SetTokenCookie(w http.ResponseWriter, token string) {
	// Create a new cookie
	cookie := &http.Cookie{
		Name:    "token",                        // Cookie name
		Value:   token,                          // Token value
		Expires: time.Now().Add(24 * time.Hour), // Expires in 24 hours                          // Cannot be accessed by JavaScript
		Path:    "/",                            // Cookie is valid for all paths
	}
	// Set the cookie
	http.SetCookie(w, cookie)
}
