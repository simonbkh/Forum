package database

import "database/sql"

func CreateTables(db *sql.DB) error {
	////`DROP TABLE IF EXISTS users;`,
	UsersTable := []string{`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
	 	username TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL,
		password TEXT NOT NULL,
		token TEXT
		);`,
		`CREATE TABLE IF NOT EXISTS post (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
    	user_id INTEGER NOT NULL,
    	content TEXT NOT NULL,
    	FOREIGN KEY(user_id) REFERENCES users(id)
		);`}

	for _, val := range UsersTable {
		_, err := db.Exec(val)
		if err != nil {
			return err
		}

	}
	return nil
}
