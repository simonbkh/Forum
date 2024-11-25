package queries

import (
	"database/sql"
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
	var cont int
	query := `SELECT COUNT(*) FROM users WHERE username = ? OR email = ?`
	err := database.Db.QueryRow(query, username, email).Scan(&cont)
	if err != nil {
		return false
	}
	return cont == 1
}

func GetHashedPass(email string) (string, error) {
	var pass string
	query := `SELECT password FROM users WHERE email = ?`
	err := database.Db.QueryRow(query, email).Scan(&pass)
	if err != nil {
		if err == sql.ErrNoRows {
			// Specific error when no user is found with the given email
			return "", fmt.Errorf("no user found with email %s", email)
		}
		// General database error
		return "", fmt.Errorf("error retrieving hashed password: %w", err)
	}
	return pass, nil
}

func Checkemail(email string) bool {
	var count int
	query := `SELECT COUNT(*) FROM users WHERE email = ?`
	err := database.Db.QueryRow(query, email).Scan(&count)
	if err != nil {
		return false
	}
	return count == 1
}

func CheckeToken(email string)string{
	var str string
	quick := `SELECT token FROM users WHERE email = ?`
	err := database.Db.QueryRow(quick, email).Scan(&str)
	if err != nil {
		return str
	}
	return str
}

func InserToken(tocken, email string) {
	add := `UPDATE users SET token = ? where email = ?`
	_, err := database.Db.Exec(add, tocken, email)
	if err != nil {
		fmt.Println(err)
	}
}
