package services

import (
	"errors"
	"net/http"
	"time"

	"forum/internal/logic/validators"
)

type Post struct {
	Title      string
	Content    string
	Categories []string
	Username   string
	Date       time.Time
}

var Posts []Post

// Post management logic

func Post_Service(w http.ResponseWriter, r *http.Request) error {
	title := r.FormValue("title")
	content := r.FormValue("post")
	categories := r.Form["category"]

	err := validators.TitleValidator(title)
	if err != nil {
		return err
	}
	if len(categories) == 0 {
		categories = append(categories, "general")
	} else if len(categories) > 3 {
		return errors.New("maximum categories to choose is 3")
	}

	err = validators.CategoriesValidator(categories)
	if err != nil {
		return err
	}
	NewPost := Post{
		Title:      title,
		Content:    content,
		Categories: categories,
	}
	Posts = append(Posts, NewPost)
	return nil
}
