## Build the Program
use -json flag to enable jsonfile storage
Note:By default json is disabled and inMemory store is enabled
### install go
https://go.dev/doc/install
### Build cmd
```
cmd : 
1.Linux   : GOOS=linux GOARCH=amd64 go build -o library-manager .
2.Windows : GOOS=windows GOARCH=amd64 go build -o library-manager.exe .
3.Mac     : GOOS=darwin GOARCH=amd64 go build -o library-manager .
```

## Run the Program

cmd :  `./library_mangement`
Interactive Menu
### output: 
#### Add a new Book:
```

Library Menu:
1. Add a new book
2. Remove a book
3. Search for a book
4. List books by genre
5. Display all books
6. Exit
Enter your choice: 1
Enter the title: the hopit
Enter the author: john doe
Enter the genre: fantasy
Added 'the hopit' by john doe (Genre: fantasy)
```
####  Remove a book:
```
Library Menu:
1. Add a new book
2. Remove a book
3. Search for a book
4. List books by genre
5. Display all books
6. Exit
Enter your choice: 2
Enter the title of the book to remove: hopit
Removed 'hopit'
```
#### List All Books
```
Library Menu:
1. Add a new book
2. Remove a book
3. Search for a book
4. List books by genre
5. Display all books
6. Exit
Enter your choice: 5

List of All Books:
1:'hopit' by john doe (Genre: fantasy)
```
#### Seach for a book
```
Library Menu:
1. Add a new book
2. Remove a book
3. Search for a book
4. List books by genre
5. Display all books
6. Exit
Enter your choice: 3
Enter the title or author to search: john
'hopit' by john doe (Genre: fantasy)
```
#### List by genre 
```
Library Menu:
1. Add a new book
2. Remove a book
3. Search for a book
4. List books by genre
5. Display all books
6. Exit
Enter your choice: 4
Enter the genre: fantasy
'hopit' by john doe (Genre: fantasy)
```

### Single line Command's works only with json Flag enabled
``` 
ADD ./library-manager add --title="To Kill a Mockingbird" --author="Harper Lee" -genre="fantasy" -json
#### Output
Added 'To Kill a Mockingbird' by Harper Lee (Genre: fantasy)

Remove ./library-manager remove --title="To Kill a Mockingbird" -json
Search ./library-manager search --term="To Kill a Mockingbird" -json
List ./library-manager list -genre="fantasy" -json
List ./library-manager list-all -genre="fantasy" -json
