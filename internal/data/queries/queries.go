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
func UpdiateSesiontoken(sessionToke, email string) error {
	statement, er := database.Db.Prepare(`UPDATE users SET sessionToke = ? WHERE email = ?`)
	if er != nil {
		return er
	}
	_, er = statement.Exec(sessionToke, email)
	if er != nil {
		return er
	}
	return nil
}

// /check this token  is it available
func IssesionidAvailable(sessionToke string) bool {
	query := `SELECT COUNT(*) FROM users WHERE sessionToke = ?`
	var cont int
	er := database.Db.QueryRow(query, sessionToke).Scan(&cont)
	if er != nil {
		fmt.Println(er)
		return false
	}

	return cont == 1
}

// /// remove token sisionid
func Removesesionid(sessionToke string) error {
	query := `UPDATE users SET sessionToke = NULL WHERE sessionToke = ?`
	_, er := database.Db.Exec(query, sessionToke)
	if er != nil {
		return er
	}
	return nil
}
