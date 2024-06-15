package cmd_test

import (
	"library_cli/cmd"
	"library_cli/library"
	"os"
	"testing"
)

func TestAddBookCmd(t *testing.T) {
	lib := library.NewLibrary()

	// Test adding a book with all flags
	os.Args = []string{"cmd", "add", "-title=Go Programming", "-author=John Doe", "-genre=Programming"}
	err := cmd.AddBookCmd(lib)
	if err != nil {
		t.Errorf("AddBookCmd failed: %v", err)
	}

	// Test adding a book with missing title
	os.Args = []string{"cmd", "add", "-author=John Doe", "-genre=Programming"}
	err = cmd.AddBookCmd(lib)
	if err == nil || err.Error() != "title, author, and genre must be provided" {
		t.Errorf("Expected error for missing title, got: %v", err)
	}

	// Test adding a book with JSON flag
	os.Args = []string{"cmd", "add", "-title=Go Programming", "-author=John Doe", "-genre=Programming", "-json=true"}
	err = cmd.AddBookCmd(lib)
	if err != nil {
		t.Errorf("AddBookCmd failed with JSON flag: %v", err)
	}
}

func TestListAllBooksCmd(t *testing.T) {
	lib := library.NewLibrary()

	// Test listing all books without JSON flag
	os.Args = []string{"cmd", "list"}
	err := cmd.ListAllBooksCmd(lib)
	if err != nil {
		t.Errorf("ListAllBooksCmd failed: %v", err)
	}

	// Test listing all books with JSON flag
	os.Args = []string{"cmd", "list", "-json=true"}
	err = cmd.ListAllBooksCmd(lib)
	if err != nil {
		t.Errorf("ListAllBooksCmd failed with JSON flag: %v", err)
	}
}

func TestListBooksByGenreCmd(t *testing.T) {
	lib := library.NewLibrary()

	// Test listing books by genre with all flags
	os.Args = []string{"cmd", "list", "-genre=Programming"}
	err := cmd.ListBooksByGenreCmd(lib)
	if err != nil {
		t.Errorf("ListBooksByGenreCmd failed: %v", err)
	}

	// Test listing books by genre with missing genre
	os.Args = []string{"cmd", "list"}
	err = cmd.ListBooksByGenreCmd(lib)
	if err == nil || err.Error() != "genre must be provided" {
		t.Errorf("Expected error for missing genre, got: %v", err)
	}

	// Test listing books by genre with JSON flag
	os.Args = []string{"cmd", "list", "-genre=Programming", "-json=true"}
	err = cmd.ListBooksByGenreCmd(lib)
	if err != nil {
		t.Errorf("ListBooksByGenreCmd failed with JSON flag: %v", err)
	}
}

func TestRemoveBookCmd(t *testing.T) {
	lib := library.NewLibrary()

	// Test removing a book with all flags
	os.Args = []string{"cmd", "remove", "-title=Go Programming"}
	err := cmd.RemoveBookCmd(lib)
	if err != nil {
		t.Errorf("RemoveBookCmd failed: %v", err)
	}

	// Test removing a book with missing title
	os.Args = []string{"cmd", "remove"}
	err = cmd.RemoveBookCmd(lib)
	if err == nil || err.Error() != "title must be provided" {
		t.Errorf("Expected error for missing title, got: %v", err)
	}

	// Test removing a book with JSON flag
	os.Args = []string{"cmd", "remove", "-title=Go Programming", "-json=true"}
	err = cmd.RemoveBookCmd(lib)
	if err != nil {
		t.Errorf("RemoveBookCmd failed with JSON flag: %v", err)
	}
}

func TestSearchBookCmd(t *testing.T) {
	lib := library.NewLibrary()

	// Test searching a book with all flags
	os.Args = []string{"cmd", "search", "-term=Go Programming"}
	err := cmd.SearchBookCmd(lib)
	if err != nil {
		t.Errorf("SearchBookCmd failed: %v", err)
	}

	// Test searching a book with missing term
	os.Args = []string{"cmd", "search"}
	err = cmd.SearchBookCmd(lib)
	if err == nil || err.Error() != "search term must be provided" {
		t.Errorf("Expected error for missing term, got: %v", err)
	}

	// Test searching a book with JSON flag
	os.Args = []string{"cmd", "search", "-term=Go Programming", "-json=true"}
	err = cmd.SearchBookCmd(lib)
	if err != nil {
		t.Errorf("SearchBookCmd failed with JSON flag: %v", err)
	}
}
