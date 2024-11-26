package queries

import (
	"database/sql"
	"errors"
	"fmt"
	"forum/internal/data/database"
	"forum/internal/data/utils"
	// "forum/internal/logic/services"
)

// // SetSessionToken sets a UUID session token for a user
// func SetSessionToken(db *sql.DB, username string) (string, error) {
// 	// Generate new UUID
// 	token := uuid.DefaultGenerator

// 	// Update the user's session token in the database
// 	query := `UPDATE users SET session_token = ? WHERE username = ?`
// 	result, err := db.Exec(query, token, username)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Check if the user was found and updated
// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		return "", err
// 	}
// 	if rowsAffected == 0 {
// 		return "", sql.ErrNoRows
// 	}

// 	return token, nil
// }

func Logged(token string) bool {
	var count int
	query := `SELECT COUNT(*) FROM sessions WHERE token = ?`
	err := database.Db.QueryRow(query, token).Scan(&count)
	if err != nil {
		return false
	}
	return count == 1
}

func InserUser(username, email, password string) error {
	statement, err := database.Db.Prepare(`INSERT INTO users (username, email, password) values (?,?,?)`)
	if err != nil {
		return err
	}
	defer statement.Close()
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
		return false
	}
	return count > 0
}

func InsertPost(post utils.Post) error {
	// p.(NewPost)
	id := 0
	err := QueryID(post.Username, &id)
	if err != nil {
		return errors.New("khona makaynch")
	}
	statement, err := database.Db.Prepare(`INSERT INTO posts (user_id ,title, content, username, created_at) values (?,?,?,?)`)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(id, post.Title, post.Content, post.Username, post.Date)
	if err != nil {
		return err
	}
	return nil
}

func QueryID(email string, id *int) error {
	var idd string
	query := `SELECT id FROM users WHERE email = ?`
	err := database.Db.QueryRow(query, email).Scan(&idd)
	if err != nil {
		//fmt.Println(username)
		return fmt.Errorf("no user found with username %s", idd)

	}
	return nil
}
func InsertSession(email, token string) error {
	var id int
	err := QueryID(email, &id)
	if err != nil {
		//handel this
		return err
	}
	statement, err := database.Db.Prepare(`INSERT INTO sessions (user_id , session_id) values (?,?)`)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(id,token)
	if err != nil {
		return err
	}
	return nil
}

// // SetSessionToken sets a UUID session token for a user
// func SetSessionToken(email, uuid string) error {

// 	// Update the user's session token in the database
// 	query := `UPDATE sessions SET session_id = ? WHERE user_id = ?`
// 	result, err := database.Db.Exec(query, uuid)
// 	if err != nil {
// 		return err
// 	}

// 	// Check if the user was found and updated
// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		return err
// 	}
// 	if rowsAffected == 0 {
// 		return sql.ErrNoRows
// 	}

// 	return nil
// }

// //  Function to validate a session token
// func ValidateSessionToken(db *sql.DB, token string) (string, error) {
// 	var username string
// 	query := `SELECT username FROM users WHERE session_token = ?`
// 	err := db.QueryRow(query, token).Scan(&username)
// 	if err != nil {
// 		return "", err
// 	}
// 	return username, nil
// }

func Logout(token string) error {
	fmt.Println([]byte(token[6:]))
	// Prepare the SQL statement to delete the token from the users table
	statement, err := database.Db.Prepare(`DELETE FROM sessions WHERE session_id = ?`)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err) // Wrap error for better context
	}
	defer statement.Close() // Ensure the statement is closed after execution

	_, err = statement.Exec(token[6:])
	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	return nil
}
