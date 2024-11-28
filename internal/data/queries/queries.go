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

func CheckeToken(email string) string {
	var str string
	quick := `SELECT token FROM users WHERE email = ?`
	err := database.Db.QueryRow(quick, email).Scan(&str)
	if err != nil {
		return str
	}
	return str
}

func Insert_OR_remove_token(tocken, email string) error {
	var qry string
	if email == "" {
		qry = `UPDATE users SET token = NULL where token = ?`
		_, err := database.Db.Exec(qry, tocken)
		if err != nil {
			return err
		}
	} else {
		qry = `UPDATE users SET token = ? where email = ?`
		_, err := database.Db.Exec(qry, tocken, email)
		if err != nil {
			return err
		}
	}
	return nil
}
type User struct {
    ID       int
    Username string
}
func Insert_Post(title, content, token string) error {
	var post_user User
	
	date := time.Now().Format(time.DateTime)
	query := `SELECT id, username FROM users WHERE token = ?`
	err := database.Db.QueryRow(query, token).Scan(&post_user.ID, &post_user.Username)
	if err != nil {
		return err
	}
	
	statement, err := database.Db.Prepare(`INSERT INTO post (title, content,date ,user, user_id) VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(title, content, date, post_user.Username, post_user.ID)
	if err != nil {
		return err
	}

	return nil
}
type Post struct {
    ID        int
    Title     string
    Content   string
    Date      string
    Username  string
    CreatedAt string
}

func GetPosts() ([]Post, error) {
	var posts []Post
	query := `SELECT id, title, content, date, user FROM post ORDER BY id DESC`
	rows, err := database.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Date, &post.Username)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}