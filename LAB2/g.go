package main

import("fmt")

func main(){
	var n int 
	fmt.Scan(&n)
	cntE := 0
	cntO := 0
	for i:= 0; i < n; i++{
		var x int
		fmt.Scan(&x)
		if x % 2 == 0{
			cntE++
		}else{
			cntO++
		}
	}
	fmt.Println(cntE, cntO)
}