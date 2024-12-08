package queries

import (
	"database/sql"
	"fmt"
	"time"

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

// // insert session in database
func Insersessions(sessionToken, email string, expiry time.Time) error {
	query := `select id from users where email = ?`
	var id int
	err := database.Db.QueryRow(query, email).Scan(&id)
	if err != nil {
		return err
	}
	statement, er := database.Db.Prepare(`INSERT INTO sessions (sessionToken, user_id, expiry) values (?,?,?)`)
	if er != nil {
		return er
	}
	_, er = statement.Exec(sessionToken, id, expiry)
	if er != nil {
		return er
	}
	return nil
}

// /// updiate session id of database
func UpdiateSesiontoken(sessionToke, email string, expiry time.Time) error {
	query := `select id from users where email = ?`
	var id int
	err := database.Db.QueryRow(query, email).Scan(&id)
	if err != nil {
		return err
	}

	query = `UPDATE sessions SET sessionToken = ?, expiry = ? WHERE user_id = ?`
	_, er := database.Db.Exec(query, sessionToke, expiry, id)
	if er != nil {
		return er
	}
	return nil
}

// /check this token  is it available
func IssesionidAvailable(sessionToke, email string) (bool, time.Time) {
	var expiry time.Time
	if email == "" {
		query := `select expiry from sessions where sessionToken = ?`
		err := database.Db.QueryRow(query, sessionToke).Scan(&expiry)
		if err != nil {
			return false, expiry
		}
		return err == nil, expiry
	} else {
		query := `select id from users where email = ?`
		var id int

		err := database.Db.QueryRow(query, email).Scan(&id)
		if err != nil {
			return false, time.Time{}
		}
		query = `select expiry from sessions where user_id = ?`
		err = database.Db.QueryRow(query, id).Scan(&expiry)
		if err != nil {
			return false, expiry
		}
		ex := expiry
		if ex.Before(time.Now()) {
			return false, ex
		}
	}
	return true, expiry
}

// /// remove token sisionid
func Removesesionid(sessionToke, email string) error {
	if email != "" {
		query := `select id from users where email = ?`
		var id int
		err := database.Db.QueryRow(query, email).Scan(&id)
		if err != nil {
			return err
		}
		query = `DELETE FROM sessions WHERE user_id = ?`
		_, err = database.Db.Exec(query, id)
		if err != nil {
			return err
		}

	} else {
		query := `DELETE FROM sessions WHERE sessionToken = ?`
		_, err := database.Db.Exec(query, sessionToke)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}
