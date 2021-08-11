package infra

type Postgre interface {
}

type Opsgenie interface {
}

type JSONPlaceholder interface {
	ListPosts() []JSONPlaceholderPost
}

type JSONPlaceholderPost struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
