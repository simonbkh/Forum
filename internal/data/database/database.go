package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

type Post struct {
	State  int // state of likes and dislikes
	Number []int // number of likes and dislikes
	Post_id int
	User_id    string
	Username   string
	Title      string
	Content    string
	Date       string
	Categories []string
}


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
