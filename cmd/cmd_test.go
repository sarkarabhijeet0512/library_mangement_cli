package cmd_test

import (
	"library_cli/cmd"
	"library_cli/library"
	"os"
	"path/filepath"
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
	// Remove the JSON file after test
	jsonFilePath := filepath.Join(".", "library.json")
	if err := os.Remove(jsonFilePath); err != nil && !os.IsNotExist(err) {
		t.Errorf("Failed to remove JSON file: %v", err)
	}
}

func TestListBooksByGenreCmd(t *testing.T) {
	lib := library.NewLibrary()
	// Test listing books by genre with all flags
	os.Args = []string{"cmd", "list", "-genre=Programming", "-json"}
	err := cmd.ListBooksByGenreCmd(lib)
	expectedErr := "no books found in genre 'Programming'"
	if err == nil || err.Error() != expectedErr {
		t.Errorf("Expected error for missing title: %v, got: %v", expectedErr, err)
	}
	os.Args = []string{"cmd", "add", "-title=Go Programming", "-author=John Doe", "-genre=Programming"}
	err = cmd.AddBookCmd(lib)
	if err != nil {
		t.Errorf("AddBookCmd failed: %v", err)
	}
	// Test listing books by genre with all flags
	os.Args = []string{"cmd", "list", "-genre=Programming"}
	err = cmd.ListBooksByGenreCmd(lib)
	if err != nil {
		t.Errorf("ListBooksByGenreCmd failed: %v", err)
	}

	// Test listing books by genre with missing genre
	os.Args = []string{"cmd", "list"}
	err = cmd.ListBooksByGenreCmd(lib)
	if err == nil || err.Error() != "genre must be provided" {
		t.Errorf("Expected error for missing genre, got: %v", err)
	}
	os.Args = []string{"cmd", "add", "-title=Go Programming", "-author=John Doe", "-genre=Programming", "-json"}
	err = cmd.AddBookCmd(lib)
	if err != nil {
		t.Errorf("AddBookCmd failed: %v", err)
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
	os.Args = []string{"cmd", "add", "-title=Go Programming", "-author=John Doe", "-genre=Programming"}
	err := cmd.AddBookCmd(lib)
	if err != nil {
		t.Errorf("AddBookCmd failed: %v", err)
	}
	// Test removing a book with all flags
	os.Args = []string{"cmd", "remove", "-title=Go Programming"}
	err = cmd.RemoveBookCmd(lib)
	if err != nil {
		t.Errorf("RemoveBookCmd failed: %v", err)
	}

	// Test removing a book with missing title
	os.Args = []string{"cmd", "remove"}
	err = cmd.RemoveBookCmd(lib)
	if err == nil || err.Error() != "title must be provided" {
		t.Errorf("Expected error for missing title, got: %v", err)
	}
	os.Args = []string{"cmd", "add", "-title=Go Programming", "-author=John Doe", "-genre=Programming", "-json"}
	err = cmd.AddBookCmd(lib)
	if err != nil {
		t.Errorf("AddBookCmd failed: %v", err)
	}
	// Test removing a book with JSON flag
	os.Args = []string{"cmd", "remove", "-title=Go Programming", "-json"}
	err = cmd.RemoveBookCmd(lib)
	if err != nil {
		t.Errorf("RemoveBookCmd failed with JSON flag: %v", err)
	}
	// Remove the JSON file after test
	jsonFilePath := filepath.Join(".", "library.json")
	if err := os.Remove(jsonFilePath); err != nil && !os.IsNotExist(err) {
		t.Errorf("Failed to remove JSON file: %v", err)
	}
}

func TestSearchBookCmd(t *testing.T) {
	lib := library.NewLibrary()
	os.Args = []string{"cmd", "add", "-title=Go Programming", "-author=John Doe", "-genre=Programming"}
	err := cmd.AddBookCmd(lib)
	if err != nil {
		t.Errorf("AddBookCmd failed: %v", err)
	}
	// Test searching a book with all flags
	os.Args = []string{"cmd", "search", "-term=Go Programming"}
	err = cmd.SearchBookCmd(lib)
	if err != nil {
		t.Errorf("SearchBookCmd failed: %v", err)
	}

	// Test searching a book with missing term
	os.Args = []string{"cmd", "search"}
	err = cmd.SearchBookCmd(lib)
	if err == nil || err.Error() != "search term must be provided" {
		t.Errorf("Expected error for missing term, got: %v", err)
	}
	os.Args = []string{"cmd", "add", "-title=Go Programming", "-author=John Doe", "-genre=Programming", "-json"}
	err = cmd.AddBookCmd(lib)
	if err != nil {
		t.Errorf("AddBookCmd failed: %v", err)
	}
	// Test searching a book with JSON flag
	os.Args = []string{"cmd", "search", "-term=Go Programming", "-json"}
	err = cmd.SearchBookCmd(lib)
	if err != nil {
		t.Errorf("SearchBookCmd failed with JSON flag: %v", err)
	}
	// Remove the JSON file after test
	jsonFilePath := filepath.Join(".", "library.json")
	if err := os.Remove(jsonFilePath); err != nil && !os.IsNotExist(err) {
		t.Errorf("Failed to remove JSON file: %v", err)
	}
}
func TestAllListBooksCmd(t *testing.T) {
	lib := library.NewLibrary()
	os.Args = []string{"cmd", "add", "-title=Go Programming", "-author=John Doe", "-genre=Programming"}
	err := cmd.AddBookCmd(lib)
	if err != nil {
		t.Errorf("AddBookCmd failed: %v", err)
	}
	// Test listing all books without JSON flag
	os.Args = []string{"cmd", "list"}
	err = cmd.ListAllBooksCmd(lib)
	if err != nil {
		t.Errorf("ListAllBooksCmd failed: %v", err)
	}
}
