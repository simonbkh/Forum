package handlers

import (
	"fmt"
	"net/http"
)

func Like(w http.ResponseWriter, r *http.Request) {
	//liked := r.URL.Query().Get("liked")
	id := r.PathValue("id")
	reaction := r.PathValue("reaction")
	fmt.Println(id, reaction)

}
