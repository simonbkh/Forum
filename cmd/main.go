package main

import (
	"net/http"

	data "forum/internal/data/database"
	"forum/internal/presentation/router"
)


func main() {
	db, err := data.Database()
	if err != nil {
		return
	}

	defer db.Close()

	serv := http.NewServeMux()

	err = router.Router(serv)
	if err != nil {
		return
	}
}
