package main

import("fmt")

func main(){
	var n int
	var l int
	var r int
	fmt.Scan(&n,&l,&r)
	arr:= make([]int,n)
	for i:= 1; i <= n; i++{
		fmt.Scan(&arr[i])
	}
	for i:= 1; i <l; i++{
		fmt.Println(arr[i])
	}
	for i:= r; i >= l; i--{
		fmt.Println(arr[i])
	}
	for i:= r+1; i <= n; i++{
		fmt.Println(arr[i])
	}
}