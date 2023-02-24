package main

import("fmt")

func main(){
	var n int
	fmt.Scan(&n)
	cnt:= 0
	for i:=0; i < n; i++{
		var x int
		fmt.Scan(&x)
		if x %10 == 7{
			cnt++
		}
	}
	fmt.Println(cnt)
}