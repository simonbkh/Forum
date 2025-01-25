package queries

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

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

func InsertPost(post utils.Post, id int) (string, error) {
	// p.(NewPost)
	// id := 0
	// err := QueryID(post.Username, &id)
	// err := GetId()

	statement, err := database.Db.Prepare(`INSERT INTO posts (user_id ,title, content, created_at) values (?,?,?,?)`)
	if err != nil {
		return "", err
	}
	defer statement.Close()
	_, err = statement.Exec(id, post.Title, post.Content, post.Date)
	if err != nil {
		return "", err
	}
	var post_id string
	err = database.Db.QueryRow(`
    	SELECT id 
    	FROM posts 
    	WHERE user_id = ? AND title = ? AND content = ? AND created_at = ?`,
		id, post.Title, post.Content, post.Date).
		Scan(&post_id)
	if err != nil {
		return "", err
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

func GetPosts() ([]utils.Post, error) {
	var posts []utils.Post
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
		var post utils.Post
		err := rows.Scan(&post.Post_id, &post.User_id, &post.Title, &post.Content, &post.Date)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		post.Categories, err = GetCategories(post.Post_id)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		// fmt.Println(post.Title)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		// post.Categories = append(post.Categories, cat)
		post.Username, err = GetUser(post.User_id)
		// fmt.Println(post.Username)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func GetCategories(post_id int) ([]string, error) {
	categories := []string{}
	statement, err := database.Db.Prepare(`SELECT category_name FROM categories WHERE posts_id = ?`)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer statement.Close()
	res, err := statement.Query(post_id)
	if err != nil {
		return nil, fmt.Errorf("failed to execute statement: %w", err)
	}
	defer res.Close()
	for res.Next() {
		var cat string
		err = res.Scan(&cat)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		categories = append(categories, cat)
	}
	// fmt.Println(categories)

	return categories, nil
}

func GetUser(uid int) (string, error) {
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

func InsertComment(post_id int, user_id int, comment string, date string) error {

	statement, err := database.Db.Prepare(`INSERT INTO comments (posts_id, user_id, comment, created_at) values (?,?,?,?)`)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, er := statement.Exec(user_id, post_id, comment, date)
	if er != nil {
		return er
	}

	return nil
}

func GetCommment(post_id int) ([]utils.Comment, error) {
	var cmt []utils.Comment
	// fmt.Println("000",id)
	statement, err := database.Db.Prepare(`SELECT * FROM comments  where posts_id = ?  ORDER BY created_at DESC`)

	if err != nil {
		return nil, err
	}

	defer statement.Close()
	rows, err := statement.Query(post_id)

	if err != nil {

		return nil, fmt.Errorf("failed to execute statement: %w", err)
	}
	defer rows.Close()

	for rows.Next() {

		var com utils.Comment
		err := rows.Scan(&com.ID, &com.Id_post, &com.Username, &com.Cont, &com.Date)

		if err != nil {

			return nil, err
		}

		id, _ := strconv.Atoi(com.Username)

		com.Username, err = GetUser(id)
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
