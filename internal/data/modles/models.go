package modles

type Data struct {
	UserStatus bool
}

var UserStatus bool

type Comment struct {
	ID       int
	Username string
	Id_post  int
	Cont     string
	Date     string
	State  int // 0 not reacted to , 1 liked , 2 disliked
	Reactions []int // first element is the number of likes , second is the number of dislikes
}
