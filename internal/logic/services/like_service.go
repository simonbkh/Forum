package services

import (
	"strings"

	"forum/internal/data/queries"
	"forum/internal/logic/utils"
)

func Rectionservice(blasa string, req utils.ReactionRequest, postID int) error {
	// var req utils.ReactionRequest
	// token := strings.Split(req.Token, )
	userID, _ := queries.Userid(strings.TrimPrefix(req.Token, "SessionToken="))

	var err error

	if req.Action == "add" {
		err = queries.AddReaction(blasa, postID, userID, req.Type)
	} else if req.Action == "remove" {
		err = queries.RemoveReaction(blasa, postID, userID)
	}
	return err
}
