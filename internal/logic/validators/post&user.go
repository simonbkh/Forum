package validators

import (
	"forum/internal/data/queries"
	"net/http"
)

func TockenPrisent(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("token")
	if err != nil || cookie.Value == "" || !queries.CheckToken_Prisent_or_not(cookie.Value){
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return false
	}
	return true
}
