package main

import "fmt"



func main(){
	a := 23
	b := 23
	c := "rushz"


	fmt.Println(a ==b)
	fmt.Println(a >=b)
	fmt.Println(a >=b)
	fmt.Println(a !=b)
	st := fmt.Sprintf("%b",byte(c[0]))
	fmt.Println(st)
}