package main

import("fmt")

func main(){
	var s string
	var t string = ""
	fmt.Scan(&s)
	for i:= len(s)-1; i >= 0; i--{
		t += string(s[i])
	}
	if s == t{
		fmt.Print("Yes")
	}else{
		fmt.Print("No")
	}
}