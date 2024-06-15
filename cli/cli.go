package cli

import (
	"flag"
	"fmt"
	"library_cli/cmd"
	"library_cli/jsonops"
	lib "library_cli/library"
	"log"
	"os"
)

func Run() {
	library := lib.NewLibrary()
	jsonFlag := flag.Bool("json", false, "Enable reading/writing inventory from/to library.json")
	flag.Parse()

	if *jsonFlag {
		var err error
		library.Books, err = jsonops.LoadLibraryFromFile()
		if err != nil {
			log.Fatalf("failed to load library from file: %v", err)
		}
	}

	if len(os.Args) <= 2 {
		lib.DisplayMenu(library, *jsonFlag)
		return
	}

	switch os.Args[1] {
	case "add":
		if err := cmd.AddBookCmd(library); err != nil {
			log.Fatalf(err.Error())
			return
		}
	case "remove":
		if err := cmd.RemoveBookCmd(library); err != nil {
			log.Fatalf(err.Error())
			return
		}
	case "search":
		if err := cmd.SearchBookCmd(library); err != nil {
			log.Fatalf(err.Error())
			return
		}
	case "list":
		if err := cmd.ListBooksByGenreCmd(library); err != nil {
			log.Fatalf(err.Error())
			return
		}
	case "list-all":
		if err := cmd.ListAllBooksCmd(library); err != nil {
			log.Fatalf(err.Error())
			return
		}
	default:
		fmt.Println("Expected 'add', 'remove', 'search', 'list', or 'list-all' subcommands")
		os.Exit(1)
	}

	if *jsonFlag {
		err := jsonops.SaveLibraryToFile(library.Books)
		if err != nil {
			log.Fatalf("failed to save library to file: %v", err)
		}
	}
}
