package bookstore

import (
	"errors"
	"fmt"
)

// Catalog represents information about a catalog.
type Catalog map[int]Book

// Book represents information about a book.
type Book struct {
	ID              int
	Title           string
	Author          string
	Copies          int
	PriceCents      int
	DiscountPercent int
	category        string
}

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}

func (b *Book) SetPriceCents(price int) error {
	if price < 0 {
		return fmt.Errorf("bad price %d (must not bee negative)", price)
	}
	b.PriceCents = price
	return nil
}

func (b *Book) SetCategory(category string) error {
	if category != "Autobiography" {
		return fmt.Errorf("cannot set %s to the category", category)
	}
	b.category = category
	return nil
}

func (b Book) Category() string {
	return b.category
}

func (c Catalog) GetBook(id int) (Book, error) {
	if book, ok := c[id]; !ok {
		return Book{}, fmt.Errorf("ID %d doesn't exist", id)
	} else {
		return book, nil
	}
}

func (b Book) NetPriceCents() int {
	diff := b.PriceCents * b.DiscountPercent / 100
	return b.PriceCents - diff

}

func (c Catalog) GetAllBooks() []Book {
	var result []Book
	for _, book := range c {
		result = append(result, book)
	}
	return result
}
