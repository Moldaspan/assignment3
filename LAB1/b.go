package main

import("fmt")

func main(){
	var n int
	var k int
	fmt.Scan(&n,&k)
	fmt.Println(n + ( k % 10 + k / 100 ))
}