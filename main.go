package main

import (
	"fmt"
	"log"

	"github.com/echocat/golang-kata-1/v1/csvreader"
	"github.com/echocat/golang-kata-1/v1/models"
)

func main() {
	// fmt.Println(welcomeMessage())
	author := models.Author{}
	book := models.Book{}
	magazine := models.Magazine{}

	resAuthor, AuthorErr := csvreader.Show(author)
	if AuthorErr != nil {
		log.Println(AuthorErr)
	}

	resBook, BookErr := csvreader.Show(book)
	if BookErr != nil {
		log.Println(BookErr)
	}

	resMagazine, MagazineErr := csvreader.Show(magazine)
	if MagazineErr != nil {
		log.Println(MagazineErr)
	}

	fmt.Println(resAuthor)
	fmt.Println(resBook)
	fmt.Println(resMagazine)
}

func welcomeMessage() string {
	return "Hello world!"
}
