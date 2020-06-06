package main

import "fmt"

var x int = 43  //gloval scope

func main(){
	x := 44 //main function scope
	fmt.Println(x)

	{
		x := 45 // inside curlebraces scope
		fmt.Println(x)

	}
	fmt.Println(x)


	a := incre()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())


}


func incre() func()int{
	var g int
	return func() int{
		g++
		return g
	}
}