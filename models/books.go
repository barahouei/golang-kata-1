package models

const (
	BooksCSVPath  string = "resources/books.csv"
	BooksCSVComma rune   = ';'
)

type Book struct {
	Title       string   `csv:"title"`
	ISBN        string   `csv:"isbn"`
	Authors     []Author `csv:"authors"`
	Description string   `csv:"description"`
}

func (a Book) CustomComma() rune {
	return BooksCSVComma
}

func (a Book) Path() string {
	return BooksCSVPath
}
