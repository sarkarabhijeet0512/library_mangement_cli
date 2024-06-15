package library

import (
	"testing"
)

func TestLibrary_AddBook(t *testing.T) {
	l := NewLibrary()
	err := l.AddBook("The Hobbit", "J.R.R. Tolkien", "Fantasy")
	if err != nil {
		t.Errorf("Error adding book: %v", err)
	}
	if len(l.Books) != 1 {
		t.Errorf("Expected 1 book, got %d", len(l.Books))
	}
	if l.Books[0].Title != "The Hobbit" {
		t.Errorf("Expected 'The Hobbit', got '%s'", l.Books[0].Title)
	}
}

func TestLibrary_RemoveBook(t *testing.T) {
	l := NewLibrary()
	l.AddBook("The Hobbit", "J.R.R. Tolkien", "Fantasy")
	err := l.RemoveBook("The Hobbit")
	if err != nil {
		t.Errorf("Error removing book: %v", err)
	}
	if len(l.Books) != 0 {
		t.Errorf("Expected 0 books, got %d", len(l.Books))
	}
	err = l.RemoveBook("Nonexistent Book")
	if err == nil {
		t.Errorf("Expected error when removing nonexistent book")
	}
}

func TestLibrary_SearchBook(t *testing.T) {
	l := NewLibrary()
	l.AddBook("The Hobbit", "J.R.R. Tolkien", "Fantasy")
	l.AddBook("The Fellowship of the Ring", "J.R.R. Tolkien", "Fantasy")

	err := l.SearchBook("hobbit")
	if err != nil {
		t.Errorf("Error searching for book: %v", err)
	}

	err = l.SearchBook("Harry Potter")
	if err == nil {
		t.Errorf("Expected error when searching for nonexistent book")
	}
}

func TestLibrary_ListBooksByGenre(t *testing.T) {
	l := NewLibrary()
	l.AddBook("The Hobbit", "J.R.R. Tolkien", "Fantasy")
	l.AddBook("The Fellowship of the Ring", "J.R.R. Tolkien", "Fantasy")
	l.AddBook("1984", "George Orwell", "Dystopian")

	err := l.ListBooksByGenre("Fantasy")
	if err != nil {
		t.Errorf("Error listing books by genre: %v", err)
	}

	err = l.ListBooksByGenre("Science Fiction")
	if err == nil {
		t.Errorf("Expected error when listing books by nonexistent genre")
	}
}

func TestLibrary_DisplayAllBooks(t *testing.T) {
	l := NewLibrary()
	l.AddBook("The Hobbit", "J.R.R. Tolkien", "Fantasy")
	l.AddBook("The Fellowship of the Ring", "J.R.R. Tolkien", "Fantasy")
	l.AddBook("1984", "George Orwell", "Dystopian")

	err := l.DisplayAllBooks()
	if err != nil {
		t.Errorf("Error displaying all books: %v", err)
	}

	l = NewLibrary()
	err = l.DisplayAllBooks()
	if err == nil {
		t.Errorf("Expected error when displaying books in empty library")
	}
}
