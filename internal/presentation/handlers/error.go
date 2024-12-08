package handlers

import (
	"bytes"
	"html/template"
	"net/http"
)

type Err struct {
	Status  int
	Message string
}

func Errore(w http.ResponseWriter, status int) {
	var buf bytes.Buffer
	temp, err := template.ParseFiles("/home/yjaouhar/Desktop/Github/Forum/internal/presentation/templates/errors/error.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	st_temp := Err{
		Status:  status,
		Message: http.StatusText(status),
	}
	er := temp.Execute(&buf, st_temp)
	if er != nil {
		Errore(w, http.StatusInternalServerError)
	}
	temp.Execute(w, st_temp)
}
