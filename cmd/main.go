package main

import (
	"fmt"
	"net/http"

	"forum/internal/logic/utils"
	"forum/internal/presentation/router"
)

func main() {
	serv := http.NewServeMux()
	err := router.Router(serv)
	if utils.IsErrors(err) {
		fmt.Println(err)
		return
	}
}
