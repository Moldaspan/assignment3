package main

import ("fmt")

func add(a int, b int) {
	fmt.Print(a + b)
}

func main() {
	var a int
	var b int
	fmt.Scan(&a, &b)
	add(a, b)
}
