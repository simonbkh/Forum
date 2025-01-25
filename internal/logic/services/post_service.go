package services

import (
	"fmt"
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
	content := r.FormValue("content")
	categories := r.Form["categories"]
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
		return err
	}
	NewPost := database.Post{
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

func UserPosts(id string) []database.Post {
	NewPosts := []database.Post{}
	for _, post := range Posts {
		if (post.User_id) == id {
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

func TimeDifference(newPosts, oldPosts []database.Post) []database.Post {
	newPosts = nil
	for _, post := range oldPosts {
		mainDate := post.Date
		duration := timeAgo(mainDate)
		NewPost := database.Post{
			Username:   post.Username,
			Title:      post.Title,
			Content:    post.Content,
			Categories: post.Categories,
			Date:       (duration),
		}
		newPosts = append(newPosts, NewPost)
	}
	return newPosts
}

func timeAgo(t string) string {
	parsedTime, err := time.Parse(time.RFC3339, t)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return "Invalid date"
	}

	duration := time.Since(parsedTime)
	duration += time.Hour
	switch {
	case duration < time.Minute:
		return fmt.Sprintf("%d Seconds ago", int(duration.Seconds()))
	case duration < time.Hour:
		return fmt.Sprintf("%d Minutes ago", int(duration.Minutes()))
	case duration < 24*time.Hour:
		return fmt.Sprintf("%d Hours ago", int(duration.Hours()))
	case duration < 30*24*time.Hour:
		return fmt.Sprintf("%d Days ago", int(duration.Hours()/24))
	case duration < 12*30*24*time.Hour:
		return fmt.Sprintf("%d Months ago", int(duration.Hours()/(24*30)))
	default:
		return fmt.Sprintf("%d Years ago", int(duration.Hours()/(24*365)))
	}
}

// func GetmyPosts(w http.ResponseWriter, r *http.Request) []database.Post {
	
// }
