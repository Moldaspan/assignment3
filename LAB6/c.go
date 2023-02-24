package main

import "fmt"


func main(){
	var n int
	fmt.Scan(&n)
	arr:= make([]int, n)
	found := false
	for i:= 0; i < n; i++{
		fmt.Scan(&arr[i])
	}
	var x int
	fmt.Scan(&x)
	for i:= 0; i < n; i++{
		if x == arr[i]{
			found = true
		}
	}
	if found == true{
		fmt.Print("YES")
	}else{
		fmt.Print("NO")
	}

}