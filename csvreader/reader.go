package csvreader

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/echocat/golang-kata-1/v1/models"
	"github.com/gocarina/gocsv"
)

type Target interface {
	Path() string
	CustomComma() rune
}

func ReadFile(path string) (*os.File, error) {
	file, fileErr := os.Open(path)
	if fileErr != nil {
		return nil, fileErr
	}
	// defer file.Close()

	return file, nil
}

func SetCustomComma(csvComma rune) {
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = csvComma

		return r
	})
}

func Show[M models.Magazine](target Target) ([]M, error) {
	targets := []M{}

	file, readErr := ReadFile(target.Path())
	if readErr != nil {
		return []M{}, readErr
	}
	defer file.Close()

	SetCustomComma(target.CustomComma())

	unmarshalErr := gocsv.Unmarshal(file, &targets)
	if unmarshalErr != nil {
		return []M{}, unmarshalErr
	}

	return targets, nil
}
