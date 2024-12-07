package utils

type Post struct {
	ID         int
	Username   string
	Title      string
	Content    string
	Date       string
	Categories []string
}

type Comment struct {
	Id_user      int
	Id_post int
	Cont string
	Date string
}
