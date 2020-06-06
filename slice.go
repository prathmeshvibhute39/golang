package main

import "fmt"

func main(){

	x := []int{1,2,3,4,5,6,7,8}
	fmt.Println(x[:])	
	fmt.Println(x[2:])
	fmt.Println(x[2:4])
	fmt.Println(x[:4])

	for i := 0; i <= 7; i++{
		fmt.Println(i, x[i])
	}
	z := []int{10,10,1,001,01}
	
	y := append(x,z...)

	fmt.Println(y) 

	//Deleting from slice
	o := []int{1,2,3,4,5,6,7,8,9,10}
	o = append(o[:1],o[4:]...)
	fmt.Println(o)
}