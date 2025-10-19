package bookstore

import "errors"

type Book struct {
	Title  string
	Author string
	Copies int
}

// This Buy function reduces the number of copies of a specific book in our slice.
func Buy(books []Book, title string) ([]Book, error) {
	for i, b := range books {
		if b.Title == title {
			if b.Copies == 0 {
				return books, errors.New("no copies left")
			}
			books[i].Copies--
			return books, nil
		}
	}
	return books, errors.New("book not found")
}

// This Stock function increases the number of copies of a specific book in or slice.
func Stock(books []Book, title string) ([]Book, error) {
	for i, b := range books {
		if b.Title == title {
			books[i].Copies++
			return books, nil
		}
	}
	return books, errors.New("book not found")
}
