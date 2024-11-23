package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"forum/internal/data/queries"
	"forum/internal/logic/utils"
)

var db *sql.DB

func Database() (*sql.DB, error) {
	var err error
	db, err = sql.Open("sqlite3", "../internal/data/database.db")
	if utils.IsErrors(err) {
		return nil, err
	}
	err = queries.CreateTables(db)
	if utils.IsErrors(err) {
		return nil, err
	}
	return db, nil
}
