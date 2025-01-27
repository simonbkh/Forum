package queries

import (
	"database/sql"
	"fmt"
	"time"

	"forum/internal/data/database"
	"forum/internal/data/modles"
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

// check if user exist or not
func IsUserExist(username, email string) bool {
	var count int
	query := `SELECT COUNT(*) FROM users WHERE username = ? OR email = ?`
	err := database.Db.QueryRow(query, username, email).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

// get hashed password in database
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

// check if email in database o“
func Checkemail(email string) bool {
	var count int
	query := `SELECT COUNT(*) FROM users WHERE email = ?`
	err := database.Db.QueryRow(query, email).Scan(&count)
	if err != nil {
		fmt.Println("error here")
		return false
	}

	return count == 1
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

func Hh(token string) (int, error) {
	// fmt.Println(token)
	var user_id int
	// fmt.Println(user_id)
	query := `SELECT user_id FROM sessions WHERE SessionToken = ?`
	err := database.Db.QueryRow(query, token).Scan(&user_id)
	if err != nil {
		return 0, err
	}
	return user_id, nil
}

func InsertPost(post database.Post, id int) (string, error) {
	// p.(NewPost)
	// id := 0
	// err := QueryID(post.Username, &id)
	// err := GetId()

	statement, err := database.Db.Prepare(`INSERT INTO posts (user_id ,title, content, created_at) values (?,?,?,?)`)
	if err != nil {
		fmt.Println(err)
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

func GetPosts() ([]database.Post, error) {
	var posts []database.Post
	statement, err := database.Db.Prepare(`SELECT * FROM posts ORDER BY created_at DESC`)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer statement.Close()

	rows, err := statement.Query()
	if err != nil {
		return nil, fmt.Errorf("failed to execute statement: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var post database.Post
		err := rows.Scan(&post.Post_id, &post.User_id, &post.Title, &post.Content, &post.Date)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		post.Categories, err = GetCategories(post.Post_id)
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
	// fmt.Println(",,,,,,,,,", posts[0].User_id)
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
	return categories, nil
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
	for res.Next() { /////?????
		err = res.Scan(&name)
		if err != nil {
			return "", fmt.Errorf("failed to scan row: %w", err)
		}
	}
	return name, nil
}

// func GetUser(uid string) (string, error) {
// 	name := ""

// 	err := database.Db.QueryRow(`SELECT username FROM users WHERE id = ?`, uid).Scan(&name)
// 	if err != nil {
// 		if err == sql.ErrNoRows {

// 			return "", fmt.Errorf("user not found")
// 		}
// 		return "", fmt.Errorf("failed to execute query: %w", err)
// 	}

//		return name, nil
//	}
//
// ///////////////////////////////////////////////////////////////////
func GetId(token string) (int, error) {
	var id int
	err := database.Db.QueryRow(`SELECT user_id FROM sessions WHERE sessionToken = ?`, token).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to execute query: %w", err)
	}
	return id, nil
}

func GetPost(post_id int) ([]database.Post, error) {
	var posts []database.Post

	statement, err := database.Db.Prepare(`SELECT * FROM posts WHERE user_id = ? ORDER BY created_at DESC`)
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
		var post database.Post
		err = res.Scan(&post.Post_id, &post.User_id, &post.Title, &post.Content, &post.Date)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		post.Categories, err = GetCategories(post.Post_id)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		post.Username, err = GetUser(post.User_id)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
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

func GetCommment(post_id int) ([]modles.Comment, error) {
	var cmt []modles.Comment

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

		var com modles.Comment
		err := rows.Scan(&com.ID, &com.Id_post, &com.Username, &com.Cont, &com.Date)

		if err != nil {

			return nil, err
		}

		com.Username, err = GetUser(com.Username)
		if err != nil {
			return nil, err
		}

		com.Date = timeAgo(com.Date)

		cmt = append(cmt, com)
	}

	return cmt, nil
}

func timeAgo(t string) string {
	parsedTime, err := time.Parse(time.RFC3339, t)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return "Invalid date"
	}

	duration := time.Since(parsedTime)
	duration += time.Hour
	switch {
	case duration < time.Minute:
		return fmt.Sprintf("%d Seconds ago", int(duration.Seconds()))
	case duration < time.Hour:
		return fmt.Sprintf("%d Minutes ago", int(duration.Minutes()))
	case duration < 24*time.Hour:
		return fmt.Sprintf("%d Hours ago", int(duration.Hours()))
	case duration < 30*24*time.Hour:
		return fmt.Sprintf("%d Days ago", int(duration.Hours()/24))
	case duration < 12*30*24*time.Hour:
		return fmt.Sprintf("%d Months ago", int(duration.Hours()/(24*30)))
	default:
		return fmt.Sprintf("%d Years ago", int(duration.Hours()/(24*365)))
	}
}
