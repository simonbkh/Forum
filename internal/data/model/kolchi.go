package model

type Post struct {
	ID         int
	Username   string
	Title      string
	Content    string
	Date       string
	Categories []string
}

type Comment struct {
	ID      int
	Username string
	Id_post int
	Cont    string
	Date    string
}
