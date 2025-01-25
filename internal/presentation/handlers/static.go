package handlers

import (
	"fmt"
	"net/http"
	"os"
)

func Static(w http.ResponseWriter, r *http.Request) {
	file := r.PathValue("file")
	// fmt.Println(r.URL)
	style := http.StripPrefix("/static/css/", http.FileServer(http.Dir("../internal/presentation/static/css")))

	_, err := os.ReadFile("../internal/presentation/static/css/" + file)
	if err != nil {
		// fmt.Println(file)
		// Error(w, http.StatusNotFound)
		return
	}

	style.ServeHTTP(w, r)
}
func JS(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--------")
	file := r.PathValue("file")
	// fmt.Println(r.URL)
	style := http.StripPrefix("/static/js/", http.FileServer(http.Dir("../internal/presentation/static/js")))

	_, err := os.ReadFile("../internal/presentation/static/js/" + file)
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
