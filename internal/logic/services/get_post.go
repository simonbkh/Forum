package services

import (
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
	// *mok,err = queries.GetCategories()
	// fmt.Println(mok)
	return nil
}
