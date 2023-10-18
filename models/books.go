package models

type Book struct {
	Title       string
	ISBN        string
	Authors     []Author
	Description string
}
