package main

import("fmt")

func main(){
	var s string
	small := 0
	capital := 0
	fmt.Scan(&s)
	for i:= 0; i < len(s); i++{
		if s[i] >= 'a' && s[i] <= 'z'{
			small++
		}else if s[i] >= 'A' && s[i] <= 'Z'{
			capital++
		}
	}
	fmt.Print(small," ", capital)
}