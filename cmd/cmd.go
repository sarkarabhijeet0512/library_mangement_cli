package cmd

import (
	"flag"
	"fmt"
	"library_cli/jsonops"
	"library_cli/library"
	"os"
)

func AddBookCmd(library *library.Library) error {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	title := addCmd.String("title", "", "Book title")
	author := addCmd.String("author", "", "Book author")
	genre := addCmd.String("genre", "", "Book genre")
	jsonFlag := addCmd.Bool("json", false, "Enable reading/writing inventory from/to library.json")
	addCmd.Parse(os.Args[2:])

	if *title == "" || *author == "" || *genre == "" {
		return fmt.Errorf("title, author, and genre must be provided")
	}
	if err := library.AddBook(*title, *author, *genre); err != nil {
		return err
	}
	if *jsonFlag {
		err := jsonops.SaveLibraryToFile(library.Books)
		if err != nil {
			return fmt.Errorf("failed to save library to file: %v", err)
		}
	}
	return nil
}

func ListAllBooksCmd(library *library.Library) error {
	jsonFlag := flag.Bool("json", false, "Enable reading/writing inventory from/to library.json")
	flag.Parse()
	if *jsonFlag {
		var err error
		library.Books, err = jsonops.LoadLibraryFromFile()
		if err != nil {
			return fmt.Errorf("failed to load library from file: %v", err)
		}
	}
	return library.DisplayAllBooks()
}

func ListBooksByGenreCmd(library *library.Library) error {
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	genre := listCmd.String("genre", "", "Book genre")
	jsonFlag := listCmd.Bool("json", false, "Enable reading/writing inventory from/to library.json")
	listCmd.Parse(os.Args[2:])

	if *genre == "" {
		return fmt.Errorf("genre must be provided")
	}
	if *jsonFlag {
		var err error
		library.Books, err = jsonops.LoadLibraryFromFile()
		if err != nil {
			return fmt.Errorf("failed to load library from file: %v", err)
		}
	}
	return library.ListBooksByGenre(*genre)
}

func RemoveBookCmd(library *library.Library) error {
	removeCmd := flag.NewFlagSet("remove", flag.ExitOnError)
	title := removeCmd.String("title", "", "Book title")
	jsonFlag := removeCmd.Bool("json", false, "Enable reading/writing inventory from/to library.json")
	removeCmd.Parse(os.Args[2:])

	if *title == "" {
		return fmt.Errorf("title must be provided")
	}
	if *jsonFlag {
		var err error
		library.Books, err = jsonops.LoadLibraryFromFile()
		if err != nil {
			return fmt.Errorf("failed to load library from file: %v", err)
		}
	}
	return library.RemoveBook(*title)
}

func SearchBookCmd(library *library.Library) error {
	searchCmd := flag.NewFlagSet("search", flag.ExitOnError)
	searchTerm := searchCmd.String("term", "", "Search term (title or author)")
	jsonFlag := searchCmd.Bool("json", false, "Enable reading/writing inventory from/to library.json")
	searchCmd.Parse(os.Args[2:])

	if *searchTerm == "" {
		return fmt.Errorf("search term must be provided")
	}
	if *jsonFlag {
		var err error
		library.Books, err = jsonops.LoadLibraryFromFile()
		if err != nil {
			return fmt.Errorf("failed to load library from file: %v", err)
		}
	}
	return library.SearchBook(*searchTerm)
}
