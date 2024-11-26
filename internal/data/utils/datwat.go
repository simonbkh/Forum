package utils

import "time"

type Post struct {
	Title      string
	Content    string
	Categories []string
	Username   string
	Date       time.Time
}
