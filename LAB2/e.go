package main

import("fmt")

func main(){
	var n int
	fmt.Scan(&n)
	var sum int
	for i:= 0; i <= n; i++{
		sum += i
	}
	fmt.Println(sum)
}