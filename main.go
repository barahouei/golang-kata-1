package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	comma         rune   = ';'
	authorsPath   string = "resources/authors.csv"
	booksPath     string = "resources/books.csv"
	magazinesPath string = "resources/magazines.csv"
)

func main() {
	var printflag bool
	var isbnFlag string

	flag.BoolVar(&printflag, "print", false, "print all csv files")
	flag.StringVar(&isbnFlag, "isbn", "", "search by isbn")

	flag.Parse()

	if printflag {
		authorsErr := printCSV(authorsPath)
		if authorsErr != nil {
			log.Println(authorsErr)
		}

		fmt.Println(strings.Repeat("*", 10))

		magazinesErr := printCSV(magazinesPath)
		if magazinesErr != nil {
			log.Println(magazinesErr)
		}

		fmt.Println(strings.Repeat("*", 10))

		booksErr := printCSV(booksPath)
		if booksErr != nil {
			log.Println(booksErr)
		}

		fmt.Println(strings.Repeat("*", 10))
	}

	if isbnFlag != "" {
		var result []string

		result, booksErr := searchByISBN(booksPath, isbnFlag)
		if booksErr != nil {
			log.Println(booksErr)
		}

		result, magazinesErr := searchByISBN(magazinesPath, isbnFlag)
		if magazinesErr != nil {
			log.Println(magazinesErr)
		}

		fmt.Println(result)
	}

}

func readCSV(path string) ([][]string, error) {
	file, fileErr := os.Open(path)
	if fileErr != nil {
		return nil, fileErr
	}

	reader := csv.NewReader(file)
	reader.Comma = comma

	records, readErr := reader.ReadAll()
	if readErr != nil {
		return nil, readErr
	}

	return records, nil
}

func printCSV(path string) error {
	records, readErr := readCSV(path)
	if readErr != nil {
		return readErr
	}

	for _, record := range records {
		fmt.Println(record)
	}

	return nil
}

func searchByISBN(path, isbn string) ([]string, error) {
	var result []string

	records, readErr := readCSV(path)
	if readErr != nil {
		return nil, readErr
	}

	for _, record := range records {
		for index, value := range record {
			if index == 1 {
				if strings.Compare(isbn, value) == 0 {
					result = append(result, record...)
				}
			}
		}
	}

	return result, nil
}
