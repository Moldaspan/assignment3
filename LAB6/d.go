package main

import ("fmt" 
		"math"
)

func hypo(a float64, b float64){
	var c float64 = math.Sqrt(a*a + b*b)
	fmt.Print(c)
}

func main(){
	var a float64
	var b float64
	fmt.Scan(&a, &b)
	hypo(a,b)
}