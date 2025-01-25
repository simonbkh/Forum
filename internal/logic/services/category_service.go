package services

import (
	"errors"
	"net/http"
	"strings"

	"forum/internal/data/database"
	"forum/internal/logic/validators"
)

func Category_Service(w http.ResponseWriter, r *http.Request) ([]database.Post, error) {
	path := strings.Split(r.URL.Path, "/")
	if len(path) != 3 {
		return nil, errors.New("invalid path")
	}
	category := path[2]
	if category == "" {
		return nil, errors.New("wrong category")
	}
	if strings.Contains(category,"-") {
		category = strings.ReplaceAll(category,"-"," ")
	}
	// if category == "all" {
	// 	return Posts, nil
	// }
	slice := []string{category}
	err := validators.CategoriesValidator(slice)
	if err != nil {
		return nil, err
	}
	newPosts := []database.Post{}
	for _, post := range Posts {
		for _, cat := range post.Categories {

			if cat == category {
				newPosts = append(newPosts, post)
			}
		}
	}
	// fmt.Println(newPosts)

	return newPosts, nil
}
