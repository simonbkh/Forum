package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"forum/internal/logic/services"
	"forum/internal/logic/utils"
)

func HandleReaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		HandleError(w, 405)
		return
	}
	blasa := r.PathValue("direction")

	if blasa != "posts" && blasa != "comments" {
		HandleError(w, 400)
		return
	}
	id := r.PathValue("id")
	postID, er := strconv.Atoi(id)
	if er != nil {
		HandleError(w, 400)
		return
	}
	var req utils.ReactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		HandleError(w, 400)
		return
	}

	if req.Type != "like" && req.Type != "dislike" {
		HandleError(w, 400)
		return
	}

	err := services.Rectionservice(blasa, req, postID)
	if err != nil {
		// http.Error(w, "BadRequest", http.StatusBadRequest)
		HandleError(w, 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "message": ""})
}
