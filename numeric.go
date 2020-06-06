package main

import (

	"fmt"
	"runtime"
)

var x int
var y float64

func main(){

	x = 24
	y = 12.0001
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}