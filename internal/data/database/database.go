package database

import (
	"database/sql"

	"forum/internal/logic/utils"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func Database() (*sql.DB, error) {
	var err error
	Db, err = sql.Open("sqlite3", "../internal/data/database.db")
	if utils.IsErrors(err) {
		return nil, err
	}
	err = CreateTables(Db)
	if utils.IsErrors(err) {
		return nil, err
	}
	return Db, nil
}
