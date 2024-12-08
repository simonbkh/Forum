package services

import (
	"forum/internal/data/queries"
	"forum/internal/data/utils"
)

func GetPosts(mok *[]utils.Post) error {
	var err error
	*mok, err = queries.GetPosts()
	if err != nil {
		return err
	}

	return nil
}