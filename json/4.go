package main

import (
	"encoding/json"
	"fmt"
)

type Target struct {
	Id    int     `json:"id"`
	Price float64 `json:"price"`
}

func main() {
	jsonString := `[{"id":1,"price":2.58},
					{"id":4,"price":7.15}]`

	targets := []Target{}

	err := json.Unmarshal([]byte(jsonString), &targets)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, t := range targets {
		fmt.Println(t.Id, "-", t.Price)
	}
}
