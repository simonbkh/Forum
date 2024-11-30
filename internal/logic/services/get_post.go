package services

import (
	"fmt"

	"forum/internal/data/queries"
	"forum/internal/data/utils"
)

func GetPosts(mok *[]utils.Post) error {
	var err error
	*mok, err = queries.GetPosts()
	// fmt.Println(*mok)
	if err != nil {
		return err
	}
	fmt.Println(mok)
	return nil
}
