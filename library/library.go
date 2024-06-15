package library

import (
	"fmt"
	"library_cli/book"
	"sort"
	"strings"
)

type Library struct {
	Books []book.Book
}

func NewLibrary() *Library {
	return &Library{
		Books: []book.Book{},
	}
}

func (l *Library) AddBook(title, author, genre string) error {
	newBook := book.Book{Title: title, Author: author, Genre: genre}
	l.Books = append(l.Books, newBook)
	fmt.Printf("Added %s\n", newBook)
	return nil
}

func (l *Library) RemoveBook(title string) error {
	for i, book := range l.Books {
		if strings.EqualFold(book.Title, title) {
			l.Books = append(l.Books[:i], l.Books[i+1:]...)
			fmt.Printf("Removed '%s'\n", title)
			return nil
		}
	}
	return fmt.Errorf("book '%s' not found", title)
}

func (l *Library) SearchBook(searchTerm string) error {
	results := []book.Book{}
	for _, book := range l.Books {
		if strings.Contains(strings.ToLower(book.Title), strings.ToLower(searchTerm)) || strings.Contains(strings.ToLower(book.Author), strings.ToLower(searchTerm)) {
			results = append(results, book)
		}
	}
	if len(results) > 0 {
		for _, book := range results {
			fmt.Println(book)
		}
	} else {
		return fmt.Errorf("no books found for '%s'", searchTerm)
	}
	return nil
}

func (l *Library) ListBooksByGenre(genre string) error {
	results := []book.Book{}
	for _, book := range l.Books {
		if strings.EqualFold(book.Genre, genre) {
			results = append(results, book)
		}
	}
	if len(results) > 0 {
		for _, book := range results {
			fmt.Println(book)
		}
	} else {
		return fmt.Errorf("No books found in genre '%s'\n", genre)
	}
	return nil
}

func (l *Library) DisplayAllBooks() error {
	if len(l.Books) > 0 {
		fmt.Println("\nList of All Books:")
		sort.Slice(l.Books, func(i, j int) bool {
			return l.Books[i].Title < l.Books[j].Title
		})
		for i, book := range l.Books {
			fmt.Println(fmt.Sprint(i+1, ":", book))
		}
	} else {
		return fmt.Errorf("No books in the inventory")
	}
	return nil
}
