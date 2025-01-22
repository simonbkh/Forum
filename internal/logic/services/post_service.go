package services

import (
	"net/http"
	"time"

	"forum/internal/data/database"
	"forum/internal/data/queries"
	"forum/internal/logic/validators"
)

type POST = database.Post

var Posts []database.Post

// Post management logic

func Post_Service(w http.ResponseWriter, r *http.Request) error {
	title := r.FormValue("title")
	content := r.FormValue("post")
	categories := r.Form["category"]
	// user_id := 0

	err := validators.TitleValidator(title)
	if err != nil {
		return err
	}
	err = validators.CategoriesValidator(categories)
	if err != nil {
		return err
	}
	user_id, err := validators.Allowed(w, r)
	if err != nil {
		// redirect or smtg
		return err
	}

	NewPost := database.Post{
		Title:      title,
		Content:    content,
		Categories: categories,
		Date:       time.Now().Format("2006-01-02 15:04:05"),
		// Username:     string(user_id),
	}

	// fmt.Println(NewPost)
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

func UserPosts(id int) []database.Post {
	NewPosts := []database.Post{}
	for _, post := range Posts {
		if post.User_id == id {
			NewPosts = append(NewPosts, post)
		}
	}
	return NewPosts
}

func GetPosts(mok *[]database.Post) error {
	var err error
	*mok, err = queries.GetPosts()
	// fmt.Println(*mok)
	if err != nil {
		return err
	}
	// *mok,err = queries.GetCategories()
	// fmt.Println(mok)
	return nil
}
