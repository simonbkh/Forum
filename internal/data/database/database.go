package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func Database() (*sql.DB, error) {
	var err error
	Db, err = sql.Open("sqlite3", "../internal/data/database.db")
	if err != nil {
		return nil, err
	}
	err = CreateTables(Db)
	if err != nil {
		return nil, err
	}
	return Db, nil
}
