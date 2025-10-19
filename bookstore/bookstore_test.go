package bookstore_test

import (
	"bookstore"
	"testing"
)

func TestBuy(t *testing.T) {
	t.Parallel()
	books := []bookstore.Book{
		{Title: "Half of a Yellow Sun", Author: "Chimamanda Adichie", Copies: 5},
		{Title: "Things fall Apart", Author: "Chinua Achebe", Copies: 2},
		{Title: "There was a Country", Author: "Chinua Achebe", Copies: 10},
	}

	want := 4
	result, err := bookstore.Buy(books, "Half of a Yellow Sun")

	for _, b := range result {
		if b.Title == "Half of a Yellow Sun" {
			got := b.Copies
			if want != got {
				t.Errorf("Error: %s, want %d, got %d", err, want, got)
			}
		}
	}
}

func TestStock(t *testing.T) {
	t.Parallel()
	books := []bookstore.Book{
		{Title: "Half of a Yellow Sun", Author: "Chimamanda Adichie", Copies: 5},
		{Title: "Things fall Apart", Author: "Chinua Achebe", Copies: 2},
		{Title: "There was a Country", Author: "Chinua Achebe", Copies: 10},
	}

	want := 3
	result, err := bookstore.Stock(books, "Things fall Apart")

	for _, b := range result {
		if b.Title == "Things fall Apart" {
			got := b.Copies
			if got != want {
				t.Errorf("Error: %s, want %d, got %d", err, want, got)
			}
		}
	}
}
