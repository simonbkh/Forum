package database

import (
	"database/sql"
)

func CreateTables(db *sql.DB) error {
	UsersTable := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
	 		username TEXT NOT NULL UNIQUE,
			email TEXT NOT NULL unique,
			password TEXT NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS posts (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id INTEGER NOT NULL,
            title TEXT NOT NULL,
            content TEXT NOT NULL,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (user_id) REFERENCES users(id)
        )`,
		`CREATE TABLE IF NOT EXISTS sessions (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			sessionToken TEXT NOT NULL UNIQUE,
			user_id INTEGER NOT NULL UNIQUE,
			expiry TIMESTAMP NOT NULL,
			FOREIGN KEY(user_id) REFERENCES users(id)
		)`,
		`CREATE TABLE IF NOT EXISTS categories (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            posts_id INTEGER NOT NULL,
            category_name TEXT,
            FOREIGN KEY (posts_id) REFERENCES posts(id)
        )`,
		`CREATE TABLE IF NOT EXISTS comments (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            posts_id INTEGER NOT NULL,
            user_id INTEGER NOT NULL,
            comment TEXT,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (posts_id) REFERENCES posts(id)
        )`,
		`CREATE TABLE IF NOT EXISTS likes (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            post_id INTEGER NOT NULL,
            user_id INTEGER NOT NULL,
            reaction TEXT NOT NULL,
            FOREIGN KEY (post_id) REFERENCES posts(id)
            FOREIGN KEY ( user_id ) REFERENCES users(id)
			UNIQUE(user_id, post_id)  

        )`,
	}
	for _, v := range UsersTable {
		_, err := db.Exec(v)
		if err != nil {
			return err
		}
	}

	return nil
}
