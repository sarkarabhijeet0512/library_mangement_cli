// book/book_test.go
package book

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestBookString(t *testing.T) {
	b := Book{
		Title:  "Test Book",
		Author: "Test Author",
		Genre:  "Test Genre",
	}
	expected := "'Test Book' by Test Author (Genre: Test Genre)"
	if got := b.String(); got != expected {
		t.Errorf("String() = %v, want %v", got, expected)
	}
}

func TestBookJSONMarshalling(t *testing.T) {
	b := Book{
		Title:  "Test Book",
		Author: "Test Author",
		Genre:  "Test Genre",
	}
	expectedJSON := `{"title":"Test Book","author":"Test Author","genre":"Test Genre"}`
	jsonBytes, err := json.Marshal(b)
	if err != nil {
		t.Fatalf("Error marshalling book to JSON: %v", err)
	}
	if string(jsonBytes) != expectedJSON {
		t.Errorf("JSON marshal result = %s, want %s", jsonBytes, expectedJSON)
	}

	var unmarshalledBook Book
	err = json.Unmarshal([]byte(expectedJSON), &unmarshalledBook)
	if err != nil {
		t.Fatalf("Error unmarshalling JSON to book: %v", err)
	}
	if !reflect.DeepEqual(b, unmarshalledBook) {
		t.Errorf("Unmarshalled book = %v, want %v", unmarshalledBook, b)
	}
}
