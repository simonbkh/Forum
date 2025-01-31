package utils

type ReactionRequest struct {
	Token string `json:"SessionToken"` 
    Type   string `json:"type"`   // "like" or "dislike"
    Action string `json:"action"` // "add" or "remove"
}