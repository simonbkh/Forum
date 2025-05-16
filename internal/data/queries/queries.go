package queries

import (
	"database/sql"
	"fmt"
	"strconv"
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

func Userid(token string) (int, error) {
	var user_id int
	query := `SELECT user_id FROM sessions WHERE SessionToken = ?`
	err := database.Db.QueryRow(query, token).Scan(&user_id)
	if err != nil {
		return 0, err
	}
	return user_id, nil
}

func InsertPost(post database.Post, id int) (string, error) {
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

func ValidPostId(postid int) bool {
	count := 0
	query := `SELECT COUNT(*) FROM posts WHERE id = ?`
	err := database.Db.QueryRow(query, postid).Scan(&count)
	if err != nil {
		return false
	}

	return count == 1
}

func ValidCommentId(id int) bool {
	count := 0
	query := `SELECT COUNT(*) FROM comments WHERE id = ?`
	err := database.Db.QueryRow(query, id).Scan(&count)
	if err != nil {
		return false
	}

	return count == 1
}

func GetPosts(token string) ([]database.Post, error) {
	var posts []database.Post
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
		var post database.Post
		err := rows.Scan(&post.Post_id, &post.User_id, &post.Title, &post.Content, &post.Date)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		post.Categories, err = GetCategories(post.Post_id)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		if modles.UserStatus {
			id, err := GetId(token)
			if err != nil {
				return nil, err
			}
			post.State, err = GetState(post.Post_id, "post", id)
			if err != nil {
				return nil, err
			}
		}
		post.Number, err = getReactions(post.Post_id, "post_id")
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

func getReactions(entityID int, entityType string) ([]int, error) {
	if entityType != "post_id" && entityType != "comment_id" {
		return nil, fmt.Errorf("invalid entity type: %s", entityType)
	}
	var counts []int
	var likeCount int
	err := database.Db.QueryRow(
		fmt.Sprintf("SELECT COUNT(*) FROM likes WHERE %s = ? AND reaction = 'like'", entityType),
		entityID,
	).Scan(&likeCount)
	if err != nil {
		return nil, fmt.Errorf("failed to query like count: %v", err)
	}

	// Query to get the count of dislikes
	var dislikeCount int
	err = database.Db.QueryRow(
		fmt.Sprintf("SELECT COUNT(*) FROM likes WHERE %s = ? AND reaction = 'dislike'", entityType),
		entityID,
	).Scan(&dislikeCount)
	if err != nil {
		return nil, fmt.Errorf("failed to query dislike count: %v", err)
	}

	// Append the counts to the slice
	counts = append(counts, likeCount, dislikeCount)

	return counts, nil
}

func GetState(entityID int, entityType string, userID int) (int, error) {
	var reaction string
	var column string

	// Determine if we're checking for a post or a comment
	if entityType == "post" {
		column = "post_id"
	} else if entityType == "comment" {
		column = "comment_id"
	} else {
		return 0, fmt.Errorf("invalid entity type: %s", entityType)
	}

	// Prepare the query based on the entity type
	statement, err := database.Db.Prepare(fmt.Sprintf(`
		SELECT reaction FROM likes WHERE user_id = ? AND %s = ?`, column))
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer statement.Close()

	// Execute the query and get the reaction for the given entity (post or comment)
	row := statement.QueryRow(userID, entityID)
	err = row.Scan(&reaction)
	if err != nil {
		if err == sql.ErrNoRows {
			// If there's no reaction, return 0 (not liked or disliked)
			return 0, nil
		}
		return 0, fmt.Errorf("failed to scan reaction: %w", err)
	}

	// Translate the reaction into state
	var state int
	switch reaction {
	case "like": // Liked
		state = 1
	case "dislike": // Disliked
		state = 2
	default: // 0 or any other value (not liked)
		state = 0
	}
	return state, nil
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

func GetCommment(post_id int, token string) ([]modles.Comment, error) {
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

		// com.Date = timeAgo(com.Date)
		if token != "" {
			id, err := GetId(token)
			if err != nil {
				return nil, err
			}
			com.State, err = GetState(com.ID, "comment", id)
			if err != nil {
				return nil, err
			}
		}
		com.Reactions, err = getReactions(com.ID, "comment_id")
		if err != nil {
			return nil, err
		}
		cmt = append(cmt, com)
	}

	return cmt, nil
}

func AddReaction(blasa string, PID, UID int, typee string) error {
	var statement *sql.Stmt
	var err error
	if !ValidPostId(PID) && blasa == "posts" || !ValidCommentId(PID) && blasa == "comments" {
		return fmt.Errorf("NOT A VALID POST")
	}

	if blasa == "posts" {
		statement, err = database.Db.Prepare(`
			INSERT INTO likes (post_id, user_id, reaction, type) 
			VALUES (?, ?, ?, ?) 
			ON CONFLICT(user_id, post_id) 
			DO UPDATE SET reaction = EXCLUDED.reaction, type = EXCLUDED.type`)
	} else if blasa == "comments" {
		statement, err = database.Db.Prepare(`
			INSERT INTO likes (comment_id, user_id, reaction, type) 
			VALUES (?, ?, ?, ?) 
			ON CONFLICT(user_id, comment_id) 
			DO UPDATE SET reaction = EXCLUDED.reaction, type = EXCLUDED.type`)
	} else {
		return fmt.Errorf("invalid entity type: %s", blasa)
	}

	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer statement.Close()

	if blasa == "posts" {
		_, err = statement.Exec(PID, UID, typee, blasa[:len(blasa)-1])
	} else {
		_, err = statement.Exec(PID, UID, typee, blasa[:len(blasa)-1])
	}

	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	return nil
}

func RemoveReaction(blasa string, PID, UID int) error {
	var statement *sql.Stmt
	var err error

	if blasa == "comments" {
		statement, err = database.Db.Prepare(`DELETE FROM likes WHERE comment_id = ? AND user_id = ?`)
	} else {
		statement, err = database.Db.Prepare(`DELETE FROM likes WHERE post_id = ? AND user_id = ?`)
	}

	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(PID, UID)
	return err
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
	err := database.Db.QueryRow(`SELECT username FROM users WHERE id = ?`, uid).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("user not found")
		}
		return "", fmt.Errorf("failed to execute query: %w", err)
	}

	return name, nil
}

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
		id, err := strconv.Atoi(post.User_id)
		if err != nil {
			return nil, err
		}
		post.State, err = GetState(post.Post_id, "post", id)
		if err != nil {
			return nil, err
		}
		post.Number, err = getReactions(post.Post_id, "post_id")
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func GetLikedPosts(userID int) ([]database.Post, error) {
	var posts []database.Post

	// Single query to fetch all liked posts with username
	query := `
        SELECT p.id, p.user_id, p.title, p.content, p.created_at, u.username
        FROM posts p
        JOIN likes l ON p.id = l.post_id
        JOIN users u ON p.user_id = u.id
        WHERE l.user_id = ?
        ORDER BY p.created_at DESC
    `

	rows, err := database.Db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var post database.Post
		// Ensure the number of columns matches the Scan variables
		if err := rows.Scan(&post.Post_id, &post.User_id, &post.Title, &post.Content, &post.Date, &post.Username); err != nil {
			return nil, fmt.Errorf("failed to scan post: %w", err)
		}

		// Fetch categories for the post
		post.Categories, err = GetCategories(post.Post_id)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch categories for post %d: %w", post.Post_id, err)
		}
		post.State, err = GetState(post.Post_id, "post", userID)
		if err != nil {
			return nil, err
		}
		post.Number, err = getReactions(post.Post_id, "post_id")
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)

	}
	return posts, nil
}
