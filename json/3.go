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

func serialize() {
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
func main() {
	byt := []byte(`{
		"name":"Ersultan",
		"age":18,
		"is_blocked":false,
		"books":[
			{
				"name":"Kopke utylgan zhalgyz",
				"year":2020
			},
			{
				"name":"Natijeli kelissoz",
				"year":2021
			}
		]
	}`)
	var dat User
	// var dat map[string]interface{}
	//Unmarshal возварщает error
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err) // паникуем если чтоөто не так
	}
	fmt.Println(dat.Books[0].Name)
	// fmt.Println(dat["books"].([]interface{})[0].(map[string]interface{})["name"])
}
