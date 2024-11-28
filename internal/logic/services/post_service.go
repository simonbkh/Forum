package services

import (
	"errors"
	"fmt"
	"forum/internal/data/queries"
	"net/http"
	"strings"
)

// Post management logic

func PostInfo(w http.ResponseWriter, r *http.Request) error {
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
	if cookie.Value == ""  || err != nil {
		return  err
	}
	 fmt.Println(cookie.Value)
	queries.Insert_Post(title, post, cookie.Value)
	return nil
}
