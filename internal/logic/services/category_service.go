package services

import (
	"errors"
	"fmt"
	"forum/internal/data/utils"
	"forum/internal/logic/validators"
	"net/http"
	"strings"
)

func Category_Service(w http.ResponseWriter, r *http.Request) ([]utils.Post, error) {

	path := strings.Split(r.URL.Path, "/")
	if len(path) != 3 {
		return nil, errors.New("invalid path")
	}
	category := path[2]
	if category == "" {
		return nil, errors.New("wrong category")
	}
	slice := []string{category}
	err := validators.CategoriesValidator(slice)
	if err != nil {
		return nil, err
	}
	newPosts := []utils.Post{}
	for _, post := range Posts {
		for _, cat := range post.Categories {
			if cat == category {
				newPosts = append(newPosts, post)
			}
		}
	}
	fmt.Println(newPosts)

	return newPosts, nil
}