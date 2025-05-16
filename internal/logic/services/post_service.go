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

func Post_Service(w http.ResponseWriter, r *http.Request) error {
	title := r.FormValue("title")
	content := r.FormValue("content")
	categories := r.Form["categories"]

	// if len(categories) == 0 {
	// 	categories = append(categories, "All")
	// }
	err := validators.TitleValidator(title)
	if err != nil {
		return err
	}
	err = validators.ValidContent(content)
	if err != nil {
		return err
	}

	err = validators.CategoriesValidator(categories)
	if err != nil {
		return err
	}
	cook , err := r.Cookie("SessionToken")
	if err != nil {
		return err
	}
	user_id, err := queries.GetId(cook.Value)
	if err != nil {
		return err
	}
	NewPost := database.Post{
		Title:      title,
		Content:    content,
		Categories: categories,
		Date:       time.Now().Format("2006-01-02 15:04:05"),
	}

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

func GetPosts(posts *[]database.Post, token string) error {
	var err error
	*posts, err = queries.GetPosts(token)
	if err != nil {
		return err
	}
	return nil
}

func TimeDifference( oldPosts []database.Post) []database.Post {
	newPosts := []database.Post{}
	for _, post := range oldPosts {
		mainDate := post.Date
		duration := timeAgo(mainDate)
		NewPost := database.Post{
			State:      post.State,
			Number:     post.Number,
			User_id:    post.User_id,
			Post_id:    post.Post_id,
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
