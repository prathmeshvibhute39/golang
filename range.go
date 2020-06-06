package main

import "fmt"

func main(){

	x := []int {1,2,3,4,5,6,7,8,9}

	fmt.Println(x[0])
 			    //i
 	for i,v := range x{

 		fmt.Println(i,v)
 	}
}