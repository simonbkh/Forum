package services

import (
	"forum/internal/data/model"
	"forum/internal/data/queries"
)

func GetPosts(mok *[]model.Post) error {
	var err error

	*mok, err = queries.GetPosts()
	if err != nil {
		return err
	}

	return nil
}
