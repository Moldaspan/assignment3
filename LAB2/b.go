package main

import("fmt")

func main(){
	var n int
	fmt.Scan(&n)
	if n%2 != 0{
		fmt.Println("Super")
	}else if n > 2 && n < 5{
		fmt.Println("Not Super")
	}else if n > 6 && n < 20{
		fmt.Println("Super")
	}else if n > 20{
		fmt.Println("Not Super")
	}

}