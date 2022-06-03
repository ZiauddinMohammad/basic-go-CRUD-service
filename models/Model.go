package models

type Book struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var Books = []Book{
	{
		ID:     1,
		Title:  "Let us c",
		Author: "BalaguruSwamy",
	},
}

type BookResponse struct {
	Data    *Book  `json:"data,omitempty"`
	Message string `json:"message"`
}
