package queries

import (
	"database/sql"
	"fmt"
	"strings"

	"forum/internal/data/database"
	"forum/internal/data/model"
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

func Logged(token string) (int, error) {
	// fmt.Println(token)
	var user_id int
	// fmt.Println(user_id)
	query := `SELECT user_id FROM sessions WHERE session_id = ?`
	err := database.Db.QueryRow(query, token).Scan(&user_id)
	if err != nil {
		return 0, err
	}
	return user_id, nil
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

func InsertPost(post model.Post, id int) (int, error) {
	// p.(NewPost)
	// id := 0
	// err := QueryID(post.Username, &id)
	// err := GetId()

	statement, err := database.Db.Prepare(`INSERT INTO posts (user_id ,title, content, created_at) values (?,?,?,?)`)
	if err != nil {
		return 0, err
	}
	defer statement.Close()
	_, err = statement.Exec(id, post.Title, post.Content, post.Date)
	if err != nil {
		return 0, err
	}
	var post_id int
	err = database.Db.QueryRow(`
    	SELECT id FROM posts WHERE user_id = ? AND title = ? AND content = ? AND created_at = ?`,
		id, post.Title, post.Content, post.Date).
		Scan(&post_id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return post_id, nil
}

func QueryID(email string, id *int) error {
	var idd int
	query := `SELECT id FROM users WHERE email = ?`
	err := database.Db.QueryRow(query, email).Scan(&idd)
	if err != nil {
		// fmt.Println(username)
		return fmt.Errorf("no user found with email %d", idd)
	}
	// if err != nil {
	// 	fmt.Println(err)
	// }

	*id = idd
	return nil
}

func InsertSession(email, token string) error {
	var id int
	err := QueryID(email, &id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// get id

	statement, err := database.Db.Prepare(`INSERT INTO sessions (user_id , session_id) values (?,?)`)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(id, token)
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
	// Prepare the SQL statement to delete the token from the users table
	statement, err := database.Db.Prepare(`DELETE FROM sessions WHERE session_id = ?`)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer statement.Close()

	_, err = statement.Exec(token[6:])
	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	return nil
}

func GetPosts() ([]model.Post, error) {
	var posts []model.Post
	statement, err := database.Db.Prepare(`SELECT * FROM posts ORDER BY created_at DESC`)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer statement.Close()

	rows, err := statement.Query()
	if err != nil {
		return nil, fmt.Errorf("failed to execute statement: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var post model.Post
		err := rows.Scan(&post.ID, &post.Username, &post.Title, &post.Content, &post.Date)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		post.Username, err = GetUser(post.Username)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func GetCommment(id int) ([]model.Comment, error) {
	var cmt []model.Comment
	statement, err := database.Db.Prepare(`SELECT * FROM comment  where posts_id = ?  ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	rows, err := statement.Query(id)
	if err != nil {
		return nil, fmt.Errorf("failed to execute statement: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var com model.Comment
		err := rows.Scan(&com.ID, &com.Id_user, &com.Id_post, &com.Cont, &com.Date)
		if err != nil {
			return nil, err
		}
		com.Date = strings.ReplaceAll(com.Date, "T", " / ")
		if len(com.Date) != 0 {
			com.Date = com.Date[:len(com.Date)-1]
		}
		cmt = append(cmt, com)
	}

	return cmt, nil
}

func GetUser(uid string) (string, error) {
	name := ""
	statement, err := database.Db.Prepare(`SELECT username FROM users WHERE id = ?`)
	if err != nil {
		return "", fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer statement.Close()
	res, err := statement.Query(uid)
	if err != nil {
		return "", fmt.Errorf("failed to execute statement: %w", err)
	}
	defer res.Close()
	for res.Next() {
		err = res.Scan(&name)
		if err != nil {
			return "", fmt.Errorf("failed to scan row: %w", err)
		}
	}
	return name, nil
}

func InsertCategories(categories []string, post_id string) error {
	statement, err := database.Db.Prepare(`INSERT INTO categories (posts_id,category_name) values (?,?)`)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer statement.Close()
	for _, cat := range categories {
		_, err = statement.Exec(post_id, cat)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetIdUser(token string) (int, error) {
	Query := `SELECT user_id FROM sessions WHERE session_id = ? `
	id := 0
	err := database.Db.QueryRow(Query, token).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func InsertComment(post int, id int, comment string, date string) error {
	statement, err := database.Db.Prepare(`INSERT INTO comment (posts_id, id_user, comment, created_at) values (?,?,?,?)`)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, er := statement.Exec(post, id, comment, date)
	if er != nil {
		return er
	}

	return nil
}
