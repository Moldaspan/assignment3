package main

import("fmt")

func main(){
	var n int
	fmt.Scan(&n)
	arr:=make([]int,n)
	for i:= 0; i < n; i++{
		fmt.Scan(&arr[i])
	}
	mx:=arr[0]
	mn:=arr[0]
	for i:= 0; i < n; i++{
		if arr[i] > mx{
			mx = arr[i]
		}
		if arr[i] < mn{
			mn = arr[i]
		}
	}
	for i:= 0; i < n; i++{
		if arr[i] == mx{
			fmt.Println(mn)
		}else{
		fmt.Println(arr[i])
		}
	}
}