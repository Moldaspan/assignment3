package main

import("fmt")

func main(){
	var n int
	fmt.Scan(&n)
	arr:= make([]int, n)
	for i:= 0; i < n; i++{
		fmt.Scan(&arr[i])
	}
	mx:=0
	for i:=0; i < n; i++{
		if arr[i] > mx{
			mx = arr[i]
		}
	}
	fmt.Println(mx)
}