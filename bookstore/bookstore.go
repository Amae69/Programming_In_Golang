package bookstore

import "errors"

type Book struct {
	Title  string
	Author string
	Copies int
}

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return b, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}

func Stock(b Book) Book {
	b.Copies++
	return b
}
