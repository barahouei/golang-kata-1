package models

type Magazine struct {
	Title       string
	ISBN        string
	Authors     []Author
	PublishedAt string
}
