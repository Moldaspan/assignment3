package main

import("fmt")

func main(){
	var n int
	fmt.Scan(&n)
	if n == 0{
		fmt.Println("None")
	}else if n%2 == 0{
		fmt.Println("Even")
	}else{
		fmt.Println("Odd")
	}
}