package bookstore

import (
	"errors"
	"fmt"
)

// Book represents information about a book.
type Book struct {
	Title  string
	Author string
	Copies int
}

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}

func GetAllBooks(catalog map[int]Book) []Book {
	var result []Book
	for _, book := range catalog {
		result = append(result, book)
	}
	return result
}

func GetBook(catalog map[int]Book, id int) (Book, error) {
	if book, ok := catalog[id]; !ok {
		return Book{}, fmt.Errorf("ID %d doesn't exist", id)
	} else {
		return book, nil
	}
}
