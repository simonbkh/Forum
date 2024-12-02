package services

import (
	"errors"
	"net/http"
	"strings"

	"forum/internal/data/queries"
	"forum/internal/logic/validators"
)

// Post management logic

func PostInfo(w http.ResponseWriter, r *http.Request) error {
	if !validators.TockenPrisent(w, r) {
		return nil
	}
	string_catigoru := "AllGeneralGamesSportsFashionTravelFoodHealth"
	title := r.FormValue("title")
	post := r.FormValue("posts")
	Categories := r.Form["Categories[]"]
	for _, category := range Categories {
		if !strings.Contains(string_catigoru, category) {
			return errors.New("invalid category")
		}
	}
	if title == "" || post == "" {
		return errors.New("title and post can't be empty")
	}
	cookie, err := r.Cookie("token")
	if cookie.Value == "" || err != nil {
		return errors.New("unauthorized")
	}
	
	queries.Insert_Post(title, post, cookie.Value, Categories)
	return nil
}
