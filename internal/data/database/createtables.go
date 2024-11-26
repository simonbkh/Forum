package database

import "database/sql"

func CreateTables(db *sql.DB) error {
	UsersTable := (`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
	 		username TEXT NOT NULL UNIQUE,
			email TEXT NOT NULL,
			password TEXT NOT NULL
		)`)

	PostTable := (`
			CREATE TABLE IF NOT EXISTS posts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			title TEXT NOT NULL,
			content TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id)
		)`)

	_, err := db.Exec(UsersTable,PostTable)
	if err != nil {
		return err
	}
	return nil
}
