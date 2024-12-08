package handlers

import (
	"net/http"
	"os"
)

func Static(w http.ResponseWriter, r *http.Request) {
	file := r.PathValue("file")
	style := http.StripPrefix("/static/css/", http.FileServer(http.Dir("../internal/presentation/static/css")))
	_, err := os.ReadFile("../internal/presentation/static/css/" + file)
	if err != nil {
		Errore(w, http.StatusNotFound)
	}

	style.ServeHTTP(w, r)
}
