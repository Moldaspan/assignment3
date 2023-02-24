package main

import("fmt")

func main(){
	var n int 
	var m int
	fmt.Scan(&n, &m)
	if n == m{
		fmt.Println("0")
	}else if n > m{
		fmt.Println("1")
	}else{
		fmt.Println("2")
	}
}