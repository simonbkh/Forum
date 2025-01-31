package services

import (
	"strings"

	"forum/internal/data/queries"
	"forum/internal/logic/utils"
)

func Rectionservice(req utils.ReactionRequest, postID int) error {
	// var req utils.ReactionRequest
	// token := strings.Split(req.Token, )
	userID, _ := queries.Hh(strings.TrimPrefix(req.Token, "SessionToken="))

	var err error

	// fmt.Println(req)
	if req.Action == "add" {
		err = queries.AddReaction(postID, userID, req.Type)
	} else if req.Action == "remove" {
		err = queries.RemoveReaction(postID, userID)
	}
	// fmt.Println(err)
	return err
}
