package main

import("fmt")

func main(){
	var n int
	fmt.Scan(&n)
	arr:= make([][]int, n)
	for i:= 0; i < n; i++{
		arr[i] = make([]int,n)
		for j :=0; j < n; j++{
			fmt.Scan(&arr[i][j])
		}
	}
	mx1:= 0
	mx2:= 0
	for i:= 0; i < n; i++{
		for j:=0; i < n; i++{
			if arr[i][j] > mx1{
				mx1 = arr[i][j]
			}
		}
		for i:= 0; i < n; i++{
			for j:=0; j < n; j++{
				if arr[i][j] > mx2 && arr[i][j] < mx1{
					mx2 = arr[i][j]
				}
			}
		}
	}
	if mx1 == mx2{
		fmt.Println(0)
	}else{
		fmt.Println(mx2)
	}
}
