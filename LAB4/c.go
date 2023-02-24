package main

import("fmt")

func main(){
	var n int
	var m int
	fmt.Scan(&n, &m)
	arr:= make([][]int, n)
	for i:= 0; i < n; i++{
		arr[i] = make([]int,m)
		for j :=0; j < n; j++{
			fmt.Scan(&arr[i][j])
		}
	}
	cnt:=0
	for i:= 0; i < n; i++{
		for j:=0; i < n; i++{
			if arr[i][j] > 0{
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}