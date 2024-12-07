package services

import (
	"errors"
	"net/http"
	"time"

	"forum/internal/data/queries"
	"forum/internal/data/utils"
	"forum/internal/logic/validators"
)

type POST = utils.Post

var Posts []utils.Post

// Post management logic

func Post_Service(w http.ResponseWriter, r *http.Request) error {
	title := r.FormValue("title")
	content := r.FormValue("post")
	categories := r.Form["category"]
	user_id := 0

	// err := validators.TitleValidator(title)
	// if err != nil {
	// 	return err
	// }
	if len(categories) == 0 {
		categories = append(categories, "general")
	} else if len(categories) > 3 {
		return errors.New("maximum categories to choose is 3")
	}

	err := validators.CategoriesValidator(categories)
	if err != nil {
		return err
	}
	user_id, err = validators.Allowed(w, r)
	if err != nil {
		// redirect or smtg
		return err
	}

	NewPost := utils.Post{
		Title:      title,
		Content:    content,
		Categories: categories,
		Date:       time.Now().Format("2006-01-02 15:04:05"),
		// Username:     string(user_id),
	}

	// Posts = append(Posts, NewPost)
	post_id, err := queries.InsertPost(NewPost, user_id)
	if err != nil {
		return err
	}
	err = queries.InsertCategories(categories, post_id)
	if err != nil {
		return err
	}
	return nil
}
