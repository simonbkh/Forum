// package handlers

// import (
// 	"fmt"
// 	"net/http"
// )

// func Like(w http.ResponseWriter, r *http.Request) {
// 	//liked := r.URL.Query().Get("liked")
// 	id := r.PathValue("id")
// 	reaction := r.PathValue("reaction")
// 	fmt.Println(id, reaction)

// }

// models/reaction.go

// database/reactions.go

// handlers/reactions.go
package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"forum/internal/logic/services"
	"forum/internal/logic/utils"
	// "forum/internal//database"
)

func HandleReaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	post_id := r.PathValue("post_id")
	// fmt.Println(post_id)
	postID, er := strconv.Atoi(post_id)
	if er != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}
	var req utils.ReactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Type != "like" && req.Type != "dislike" {
		http.Error(w, "Invalid reaction type", http.StatusBadRequest)
		return
	}

	err := services.Rectionservice(req, postID)
	if err != nil {
		http.Error(w, "BadRequest", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "message": ""})
	//  reaction := logic.Reaction{Type: req.Type, PostID: postID}

	// Handle the reaction
	// var err error

	// if req.Action == "add" {
	//     err = database.AddReaction(db, postID, userID, req.Type)
	// } else if req.Action == "remove" {
	//     err = database.RemoveReaction(db, postID, userID, req.Type)
	// } else {
	//     http.Error(w, "Invalid action", http.StatusBadRequest)
	//     return
	// }

	// if err != nil {
	//     http.Error(w, "Internal server error", http.StatusInternalServerError)
	//     return
	// }

	// Get updated counts
	// likes, dislikes, err := database.GetReactionCounts(db, postID)
	// if err != nil {
	//     http.Error(w, "Error getting updated counts", http.StatusInternalServerError)
	//     return
	// }

	// Return updated counts
	// response := struct {
	//     Likes    int `json:"likes"`
	//     Dislikes int `json:"dislikes"`
	// }{
	//     Likes:    likes,
	//     Dislikes: dislikes,
	// }
	// w.WriteHeader(200)

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(response)
}
