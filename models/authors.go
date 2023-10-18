package models

const (
	AuthorsCSVPath string = "resources/authors.csv"
	AuthorCSVComma rune   = ';'
)

type Author struct {
	Email     string `csv:"email"`
	FirstName string `csv:"firstname"`
	LastName  string `csv:"lastname"`
}

func (a Author) CustomComma() rune {
	return AuthorCSVComma
}

func (a Author) Path() string {
	return AuthorsCSVPath
}
