package main

import("fmt")

func main(){
	var n int 
	fmt.Scan(&n)
	cnt:=0
	for i:= 0; i < n; i++{
		var x int 
		fmt.Scan(&x)
		if x == 0{
			cnt++
		}
		for j:= x; j != 0; j/=10 {
			if j%10 == 0{
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}