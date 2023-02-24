package main

import (
	"fmt"
	"math"
)

func diff(n int, a[]int, b[]int ){
	for i:= 0; i < n; i++{
		fmt.Print(math.Abs(float64(a[i] - b[i])))
	}
}

func main(){
	var n int
	fmt.Scan(&n)
	a := make([]int,n)
	b:= make([]int,n)
	for i:= 0; i < n; i++{
		fmt.Scan(&a[i])
	}
	for i:=0; i < n; i++{
		fmt.Scan(&b[i])
	}
	diff(n ,a, b)
}