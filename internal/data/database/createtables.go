package database

import "database/sql"

func CreateTables(db *sql.DB) error {
	UsersTable := []string{`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
	 		username TEXT NOT NULL UNIQUE,
			email TEXT NOT NULL,
			password TEXT NOT NULL,
			sessionToke TEXT UNIQUE
		)`,
		`CREATE TABLE IF NOT EXISTS post (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
	 		titele TEXT NOT NULL,
			content TEXT NOT NULL,
			categorie TEXT NOT NULL,
			user_id INTEGER NOT NULL,			
			FOREIGN KEY(user_id) REFERENCES users(id)
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
