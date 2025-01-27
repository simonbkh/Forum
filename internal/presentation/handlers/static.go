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

func Image(w http.ResponseWriter, r *http.Request) {
	file := r.PathValue("file")

	// Specify the directory where images are stored
	imageDir := "../internal/presentation/static/images/"

	// Attempt to read the requested image file
	_, err := os.ReadFile(imageDir + file)
	if err != nil {
		// If the image does not exist, respond with a 404 error
		http.Error(w, "Image not found", http.StatusNotFound)
		return
	}

	// Serve the image file using FileServer
	imageServer := http.StripPrefix("/static/images/", http.FileServer(http.Dir(imageDir)))
	imageServer.ServeHTTP(w, r)
}
