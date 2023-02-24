package main

import("fmt")

func main(){
	var n int
	var pos_row = 1
	var pos_col = 1
	fmt.Scan(&n)
	arr:= make([][]int,n+1)
	for i:= 1; i <= n; i++{
		arr[i] = make([]int,n+1)
		for j:=1; j <= n; j++{
			fmt.Scan(&arr[i][j])
		}
	}
	mx:=arr[1][1]
	for i:= 1; i <= n; i++{
		for  j:= 1; j <= n; j++{
			if arr[i][j] > mx{
				mx = arr[i][j]
				pos_row = i
				pos_col = j
			}
		}
	}
	fmt.Print(pos_row, " ", pos_col)
}
