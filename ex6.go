package main

import "fmt"
const (

	_ = iota
	y = 2020 - iota
	z
	q
	x
)
func main (){

	fmt.Println(y,z,q,x)

}