package main

import (
   "fmt"
)
func main(){
	var n  int      
	var mx int = -1e9 
   	fmt.Scan(&n)
   	arr := make([][]int, n)
   	for i := 0; i < n; i++ {
    	arr[i] = make([]int, n)
		for j := 0; j < n; j++ {
        	fmt.Scan(&arr[i][j])      
		}
   	} 
   	for i := 0; i < n; i++ {
    	for j := 0; j < n; j++ {        
			if arr[i][j] > mx {
            	mx = arr[i][j]         
			}
      	}   
	}
   	fmt.Println(mx)
}
