package main

import("fmt")

func main(){
	var n int
	fmt.Scan(&n)
	var mx int
	mx = -1e9
	for i:= 1; i <= n; i++{
		var x int
		fmt.Scan(&x)
		if x > mx {
			mx = x
		}
	}
	fmt.Println(mx)
}