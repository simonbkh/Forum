package handlers

import (
	"fmt"
	"net/http"
	"os"
)

func Static(w http.ResponseWriter, r *http.Request) {
	file := r.PathValue("file")
	fmt.Println(file)

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



func Jss(w http.ResponseWriter, r *http.Request){
	js := r.PathValue("js")
	fmt.Println(js)

	style := http.StripPrefix("/static/js/", http.FileServer(http.Dir("../internal/presentation/static/js")))

	//	fmt.Println(style)
	// Check if the requested file exists by trying to read it
	_, err := os.ReadFile("../internal/presentation/static/js/" + js)
	if err != nil {
		// fmt.Println(file)
		// Error(w, http.StatusNotFound)
		return
	}

	style.ServeHTTP(w, r)
}
