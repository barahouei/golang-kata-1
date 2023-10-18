package models

const (
	MagazinesCSVPath  string = "resources/magazines.csv"
	MagazinesCSVComma rune   = ';'
)

type Magazine struct {
	Title       string   `csv:"title"`
	ISBN        string   `csv:"isbn"`
	Authors     []Author `csv:"authors"`
	PublishedAt string   `csv:"publishedAt"`
}

func (a Magazine) CustomComma() rune {
	return MagazinesCSVComma
}

func (a Magazine) Path() string {
	return MagazinesCSVPath
}
