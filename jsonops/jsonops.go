package jsonops

import (
	"encoding/json"
	"fmt"
	"library_cli/book"
	"os"
)

const jsonFileName = "library.json"

func LoadLibraryFromFile() ([]book.Book, error) {
	var books []book.Book
	data, err := os.ReadFile(jsonFileName)
	if err != nil {
		if os.IsNotExist(err) {
			return books, nil // If file does not exist, return empty list
		}
		return books, err
	}
	err = json.Unmarshal(data, &books)
	return books, err
}

func SaveLibraryToFile(books []book.Book) error {
	var existingBooks []book.Book

	// Check if the file exists
	if _, err := os.Stat(jsonFileName); err == nil {
		// Read existing contents
		data, err := os.ReadFile(jsonFileName)
		if err != nil {
			return fmt.Errorf("failed to read existing file: %v", err)
		}

		// Unmarshal existing data
		if err := json.Unmarshal(data, &existingBooks); err != nil {
			return fmt.Errorf("failed to unmarshal existing data: %v", err)
		}
	}

	// Append new books to existing books
	existingBooks = append(existingBooks, books...)

	// Marshal updated books to JSON
	data, err := json.MarshalIndent(existingBooks, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal books: %v", err)
	}

	// Write updated data back to file
	if err := os.WriteFile(jsonFileName, data, 0644); err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}

	fmt.Printf("Saved %d books to %s\n", len(books), jsonFileName)
	return nil
}
