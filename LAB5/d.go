package main

import("fmt")

func main(){
	var s string
	fmt.Scan(&s)
	var odd int = 0
	var even int = 0
	for i:=0; i < len(s); i++{
		if i%2 == 0{
			even += int(s[i] - 48)
		}else{
			odd += int(s[i] - 48)
		}
	}
	if odd == even{
		fmt.Print("YES")
	}else{
		fmt.Print(("NO"))
	}
}