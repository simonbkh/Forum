package database

import (
	"database/sql"
	"fmt"
)

func CreateTables(db *sql.DB) error {
	UsersTable := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
	 		username TEXT NOT NULL UNIQUE,
			email TEXT NOT NULL unique,
			password TEXT NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS post (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
	 		titele TEXT NOT NULL,
			content TEXT NOT NULL,
			categorie TEXT NOT NULL,
			user_id INTEGER NOT NULL,			
			FOREIGN KEY(user_id) REFERENCES users(id)
			)`,
		`CREATE TABLE IF NOT EXISTS sessions (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			sessionToken TEXT NOT NULL UNIQUE,
			user_id INTEGER NOT NULL,
			expiry TIMESTAMP NOT NULL,
			FOREIGN KEY(user_id) REFERENCES users(id)
			)`,
	}
	for _, v := range UsersTable {
		_, err := db.Exec(v)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}
