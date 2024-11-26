package services

import (
	"errors"
	"net/http"
	"time"
	"forum/internal/data/utils"
	"forum/internal/data/queries"
	"forum/internal/logic/validators"
)



var Posts []utils.Post

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
	NewPost := utils.Post{
		Title:      title,
		Content:    content,
		Categories: categories,
		Date:       time.Now().Format("2006-01-02 15:04:05"),
	}
	Posts = append(Posts, NewPost)
	err = queries.InsertPost(NewPost)
	if err != nil {
		return err
	}
	return nil
}
