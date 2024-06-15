package library

import (
	"bufio"
	"fmt"
	"library_cli/jsonops"
	"log"
	"os"
)

func DisplayMenu(library *Library, jsonFlag bool) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nLibrary Menu:")
		fmt.Println("1. Add a new book")
		fmt.Println("2. Remove a book")
		fmt.Println("3. Search for a book")
		fmt.Println("4. List books by genre")
		fmt.Println("5. Display all books")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter the title: ")
			scanner.Scan()
			title := scanner.Text()
			fmt.Print("Enter the author: ")
			scanner.Scan()
			author := scanner.Text()
			fmt.Print("Enter the genre: ")
			scanner.Scan()
			genre := scanner.Text()
			library.AddBook(title, author, genre)
		case 2:
			fmt.Print("Enter the title of the book to remove: ")
			scanner.Scan()
			title := scanner.Text()
			library.RemoveBook(title)
		case 3:
			fmt.Print("Enter the title or author to search: ")
			scanner.Scan()
			searchTerm := scanner.Text()
			library.SearchBook(searchTerm)
		case 4:
			fmt.Print("Enter the genre: ")
			scanner.Scan()
			genre := scanner.Text()
			library.ListBooksByGenre(genre)
		case 5:
			library.DisplayAllBooks()
		case 6:
			if jsonFlag {
				err := jsonops.SaveLibraryToFile(library.Books)
				if err != nil {
					log.Fatalf("failed to save library to file: %v", err)
				}
			}
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
