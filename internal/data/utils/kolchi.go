package utils

type Post struct {
	Post_id    int
	User_id    int
	Username   string
	Title      string
	Content    string
	Date       string
	Categories []string
}
type Comment struct {
	ID       int
	Username string
	Id_post  int
	Cont     string
	Date     string
}
