package queries

import (
	"fmt"

	"forum/internal/data/database"
)

func InserUser(username, email, password string) error {
	statement, err := database.Db.Prepare(`INSERT INTO users (username, email, password) values (?,?,?)`)
	if err != nil {
		return err
	}
	_, err = statement.Exec(username, email, password)
	if err != nil {
		return err
	}
	return nil
}

func IsUserExist(username, email string) bool {
    var count int
    query := `SELECT COUNT(*) FROM users WHERE username = ? OR email = ?`
    err := database.Db.QueryRow(query, username, email).Scan(&count)
    if err != nil {
        fmt.Println("Error checking user existence:", err)
        return false
    }
    return count > 0
}
