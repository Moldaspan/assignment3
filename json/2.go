package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	IsBlocked bool   `json:"is_blocked"	`
	Books     []Book `json:"books"`
}
type Book struct {
	Name string `json:"name"`
	Year int    `json:"year"`
}

func main() {
	var books []Book
	book1 := Book{
		Name: "Kopke utylgan zhalgyz",
		Year: 2020,
	}
	book2 := Book{
		Name: "Natijeli kelissoz",
		Year: 2021,
	}
	books = append(books, book1, book2)
	arr := User{
		Name:      "Ersultan",
		Age:       18,
		IsBlocked: false,
		Books:     books,
	}
	boolVar, _ := json.Marshal(arr)
	fmt.Println(string(boolVar))
}
