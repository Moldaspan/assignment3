package main

import("fmt")

func main(){
	var n int
	num:= 1
	fmt.Scan(&n)
	arr:= make([][]int,n)
	for i:= 0; i < n; i++{
		arr[i] = make([]int,n)
		for j:= 0; j < n; j++{
			if j == n - num{
				fmt.Print(num)
			}else{
				fmt.Print(".")
			}
		}
		fmt.Println()
		num++
	}
}