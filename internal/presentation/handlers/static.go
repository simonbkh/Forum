package handlers

import (
	"net/http"
	"os"
)

func Static(w http.ResponseWriter, r *http.Request) {
	file := r.PathValue("file")
	folder := r.PathValue("folder")
	if folder != "css" && folder != "js" && folder != "images" {
		HandleError(w, 404)
		return
	}
	style := http.StripPrefix("/static/"+folder+"/", http.FileServer(http.Dir("../internal/presentation/static/"+folder+"/")))
	_, err := os.ReadFile("../internal/presentation/static/" + folder + "/" + file)
	if err != nil {
		HandleError(w, 404)
		return
	}
	style.ServeHTTP(w, r)
}

