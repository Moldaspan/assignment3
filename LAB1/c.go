package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var a int = n % 2
	n = n / 2
	var b int = n % 2
	n = n / 2
	var c int = n % 2
	n = n / 2
	var d = n % 2
	fmt.Println(d*1 + c*2 + b*4 + a*8)
}