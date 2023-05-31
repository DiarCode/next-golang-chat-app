package models

type CreatePostJson struct {
	Title    string `json:"title"`
	Body     string `json:"body"`
	AuthorId int    `json:"authorId"`
}
