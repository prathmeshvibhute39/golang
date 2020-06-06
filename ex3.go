package main

import "fmt"

const (

	a int = 21
	b string  = "rushz"
	c bool = true
)

const (

	x = 24
	y = "root"
	z = false
)
func main(){

	fmt.Println(a,b,c,x,y,z)
	fmt.Printf("%T\n",a)
	fmt.Printf("%T\n",x)

}