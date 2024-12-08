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
	Id_user int
	Id_post int
	Cont    string
	Date    string
}
