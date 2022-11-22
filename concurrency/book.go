package main

import "fmt"

type Book struct {
	ID     int
	Title  string
	Author string
}

func (b *Book) String() string {
	return fmt.Sprintf("Title:\t\t%q\nAuthor:\t\t%q\n\n", b.Title, b.Author)
}

var books = []Book{
	Book{
		ID:     1,
		Title:  "Book1",
		Author: "Author1",
	},
	Book{
		ID:     2,
		Title:  "Book2",
		Author: "Author2",
	},
	Book{
		ID:     3,
		Title:  "Book3",
		Author: "Author3",
	},
	Book{
		ID:     4,
		Title:  "Book4",
		Author: "Author4",
	},
	Book{
		ID:     5,
		Title:  "Book5",
		Author: "Author5",
	},
	Book{
		ID:     6,
		Title:  "Book6",
		Author: "Author6",
	},
	Book{
		ID:     7,
		Title:  "Book7",
		Author: "Author7",
	},
	Book{
		ID:     1,
		Title:  "Book8",
		Author: "Author8",
	},
}
