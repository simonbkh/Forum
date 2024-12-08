package services

import (
	"strconv"
	"strings"

	"forum/internal/data/model"
)

// var Cmt []model.Comment

type strct struct {
	Cmts []model.Comment
	Post model.Post
}

func Comment_service(path string) (strct, error) {
	var Cmt []model.Comment
	var nill strct
	id, er := strconv.Atoi(path)
	if er != nil || id < 0 || id > len(Posts) {
		return nill, er
	}
	err := GetCommment(&Cmt, id)
	if err != nil {
		return nill, err
	}
	Comment := strct{
		Cmts: Cmt,
		Post: Posts[len(Posts)-id],
	}
	Comment.Post.Date = strings.ReplaceAll(Comment.Post.Date, "T", " / ")
	if len(Comment.Post.Date) != 0 {
		Comment.Post.Date = Comment.Post.Date[:len(Comment.Post.Date)-1]
	}
	return Comment, nil
}
