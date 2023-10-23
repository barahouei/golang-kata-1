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
	flag.BoolVar(&printflag, "print", false, "print all csv files")
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
}

func printCSV(path string) error {
	file, fileErr := os.Open(path)
	if fileErr != nil {
		return fileErr
	}

	reader := csv.NewReader(file)
	reader.Comma = comma

	records, readErr := reader.ReadAll()
	if readErr != nil {
		return readErr
	}

	for _, record := range records {
		fmt.Println(record)
	}

	return nil
}
