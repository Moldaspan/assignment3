package main

import("fmt")

func main(){
	var n int
	fmt.Scan(&n)
	cnt:=0
	arr:= make([]int,n)
	for i:=0; i < n; i++{
		fmt.Scan(&arr[i])
	}
	for i:= 0; i < n; i++{
		if arr[i] > 0{
			cnt++
		}
	}
	fmt.Println(cnt)
}