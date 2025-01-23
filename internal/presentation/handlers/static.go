package handlers

import (
	"net/http"
	"os"
)

func Static(w http.ResponseWriter, r *http.Request) {
	file := r.PathValue("file")
	// fmt.Println(r.URL)
	style := http.StripPrefix("/static/css/", http.FileServer(http.Dir("../internal/presentation/static/css")))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../internal/presentation/static"))))

	//	fmt.Println(style)
	// Check if the requested file exists by trying to read it
	_, err := os.ReadFile("../internal/presentation/static/css/" + file)
	if err != nil {
		// fmt.Println(file)
		// Error(w, http.StatusNotFound)
		return
	}

	style.ServeHTTP(w, r)
}

func StaticCat(w http.ResponseWriter, r *http.Request) {
	file := r.PathValue("file")

	style := http.StripPrefix("/category/static/css/", http.FileServer(http.Dir("../internal/presentation/static/css")))

	_, err := os.ReadFile("../internal/presentation/static/css/" + file)
	if err != nil {
		return
	}

	style.ServeHTTP(w, r)
}
