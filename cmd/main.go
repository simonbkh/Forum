package main

import (
	"net/http"

	"forum/internal/logic/utils"
	"forum/internal/presentation/router"
)

func main() {
	serv := http.NewServeMux()
	err := router.Router(serv)
	if utils.IsErrors(err) {
		return
	}
}
