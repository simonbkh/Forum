package database

import "database/sql"

func CreateTables(db *sql.DB) error{
	UsersTable := (`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
	 		username TEXT NOT NULL UNIQUE,
			email TEXT NOT NULL,
			password TEXT NOT NULL
		)`)
	
	_ ,err:= db.Exec(UsersTable)
	if err != nil {
		return err
	}
	return nil
}
