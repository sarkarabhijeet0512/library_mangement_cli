package book

import (
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
}

func (b Book) String() string {
	return fmt.Sprintf("'%s' by %s (Genre: %s)", b.Title, b.Author, b.Genre)
}
