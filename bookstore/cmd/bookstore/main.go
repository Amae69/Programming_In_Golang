package main

import (
	"bookstore"
	"fmt"
)

func main() {
	books := []bookstore.Book{
		{Title: "Things Fall Apart", Author: "Chinua Achebe", Copies: 20},
		{Title: "Half of a Yellow Sun", Author: "Chimamanda Adichie", Copies: 10},
		{Title: "There was a Country", Author: "Chinua Achebe", Copies: 5},
	}

	fmt.Println(books)
}
