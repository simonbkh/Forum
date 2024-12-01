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

func InserSisionToken(sessionToke string) error {
	statement, err := database.Db.Prepare(`INSERT INTO users (sessionToke) values (?)`)
	if err != nil {
		return err
	}
	_, err = statement.Exec(sessionToke)
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
		return false
	}
	return count > 0
}

func GetHashedPass(email string) (string, error) {
	var pass string

	query := `SELECT password FROM users WHERE email = ?`
	err := database.Db.QueryRow(query, email).Scan(&pass)
	fmt.Println(pass)
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
		fmt.Println("error here")
		return false
	}
	if count < 0 {
		return false
	}
	return true
}

// //// check is sision id is exit or no
func IssissiontokenExit(email string) bool {
	query := `SELECT sessionToke FROM users WHERE email = ?`
	var token string
	er := database.Db.QueryRow(query, email).Scan(&token)
	if er != nil {
		return false
	}
	return token != ""
}

// /// updiate session id of database
func UpdiateSesiontoken(newtoken, email string) error {
	query := `UPDATE users SET sessionToken = ? WHERE email = ?`
	_, er := database.Db.Exec(query, newtoken, email)
	if er != nil {
		return er
	}
	return nil
}

// /check this token  is it available
func IssesionidAvailable(sessionToke string) bool {
	query := `SELECT sessionToke FROM users WHERE sessionToke = ?`
	var cont int
	er := database.Db.QueryRow(query, sessionToke).Scan(&cont)
	if er != nil {
		return false
	}
	if cont <= 0 {
		return false
	}
	return true
}

// /// remove token sisionid
func Removesesionid(sessionToken string) error {
	query := `UPDATE users SET sessionToken = '' WHERE sessionToken = ?`
	_, er := database.Db.Exec(query, sessionToken)
	if er != nil {
		return er
	}
	return nil
}
