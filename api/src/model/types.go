package model

type Book struct {
	BookId int64  `json:"id" db:"book_id"`
	Title  string `json:"title" db:"title"`
	Author string `json:"author" db:"author"`
}
