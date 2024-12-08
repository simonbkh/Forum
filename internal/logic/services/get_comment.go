package services

import (
	"fmt"

	"forum/internal/data/model"
	"forum/internal/data/queries"
)

func GetCommment(Cmt *[]model.Comment, id int) error {
	var err error
	*Cmt, err = queries.GetCommment(id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
