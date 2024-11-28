package services

import (
	"errors"
	"forum/internal/data/queries"
	"forum/internal/data/utils"
	"forum/internal/logic/validators"
	"net/http"
	"time"
)

type POST = utils.Post

var Posts []utils.Post

// Post management logic

func Post_Service(r *http.Request) error {
	title := r.FormValue("title")
	content := r.FormValue("post")
	categories := r.Form["category"]
	user_id := 0

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
	err, user_id = validators.Allowed(r)
	//fmt.Println(user_id)
	if err != nil {
		//redirect or smtg
		return err
	}

	NewPost := utils.Post{
		Title:      title,
		Content:    content,
		Categories: categories,
		Date:       time.Now().Format("2006-01-02 15:04:05"),
		//Username:     string(user_id),
	}

	//Posts = append(Posts, NewPost)
	err = queries.InsertPost(NewPost, user_id)
	if err != nil {
		return err
	}
	return nil
}
