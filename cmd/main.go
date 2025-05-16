package main

import (
	"fmt"
	"net/http"
	"os"

	data "forum/internal/data/database"
	"forum/internal/presentation/router"
)

func main() {
	db, err := data.Database()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error for creating database : %v", err)
		return
	}

	defer db.Close()

	serv := http.NewServeMux()

	err = router.Router(serv)
	
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error setting up router: %v", err)
		return
	}
}
