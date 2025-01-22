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
	}
	for _, v := range UsersTable {
		_, err := db.Exec(v)
		if err != nil {
			return err
		}
	}

	return nil
}
