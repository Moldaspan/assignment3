package main

import("fmt")

func main(){
	var n int
	var l int
	var r int
	fmt.Scan(&n,&l,&r)
	sum:=0
	arr:= make([]int,n)
	for i:= 1; i <= n; i++{
		fmt.Scan(&arr[i])
	}
	for i:= l; i <= r; i++{
		sum += arr[i]
	}
	fmt.Println(sum)
}