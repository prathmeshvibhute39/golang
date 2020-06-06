package main

import "fmt"

func main(){

	var x [5] int
	fmt.Println(x)

	x[3] =4
	fmt.Println(x)
	fmt.Printf("%T\n",x)
	fmt.Println(len(x))


	y := [3] int {3,4,5}
	fmt.Println(y)

	z := [] int {5,5,5,7,6,5,8,7}
	fmt.Println(z)
}