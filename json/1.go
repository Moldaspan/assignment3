package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	boolVar, _ := json.Marshal("something")
	fmt.Println(string(boolVar))
}
