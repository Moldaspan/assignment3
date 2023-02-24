package main

import("fmt")

func main(){
	var n int
	fmt.Scan(&n)
	arr:= make([][]int, n)
	for i:= 0; i < n; i++{
		arr[i] = make([]int,n)
		for j :=0; j < n; j++{
			if i == 0 || j == 0{
				arr[i][j] = i+j
			}else{
				arr[i][j] = i*j
			}
		}
	}
	for i:=0; i < n; i++{
		for j:= 0; j < n; j++{
			fmt.Println(arr[i][j])
		}
	}
	
}

