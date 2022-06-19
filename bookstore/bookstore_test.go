package bookstore_test

import (
	"bookstore"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	_ "github.com/google/go-cmp/cmp"
)

func TestBook(t *testing.T) {
	t.Parallel()

	_ = bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()

	var b = bookstore.Book{
		Title:  "Nicholas Chuckleby",
		Author: "Charles Dickens",
		Copies: 8,
	}
	want := 7
	result, err := bookstore.Buy(b)
	if err != nil {
		t.Fatal(err)
	}
	got := result.Copies

	if got != want {
		t.Errorf("want %d copies after buying 1 copy from a stock of %d, got %d", want, b.Copies, got)
	}
}

func TestBuyErrorsIfNoCopiesLeft(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 0,
	}
	_, err := bookstore.Buy(b)
	if err == nil {
		t.Error("want error buying book when zero copies left, but got nil")
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	catalog := map[int]bookstore.Book{
		1: {Title: "For the Love of Go"},
		2: {Title: "The Power of Go: Tools"},
	}
	want := []bookstore.Book{
		{Title: "For the Love of Go"},
		{Title: "The Power of Go: Tools"},
	}

	got := bookstore.GetAllBooks(catalog)
	sort.Slice(got, func(i, j int) bool {
		return got[i].Title < got[j].Title
	})

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()
	catalog := map[int]bookstore.Book{
		1: bookstore.Book{Title: "The power of Go: Tools"},
		2: bookstore.Book{Title: "For the Love of Go"},
	}
	want := bookstore.Book{
		Title: "For the Love of Go",
	}
	got, _ := bookstore.GetBook(catalog, 2)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBookBadIDReturnsError(t *testing.T) {
	catalog := map[int]bookstore.Book{}
	_, err := bookstore.GetBook(catalog, 999)
	if err == nil {
		t.Fatal("want error for non-existent ID, got nil")
	}
}
