package main

import "fmt"

func main(){

	// x := 24
	
	// fmt.Println(&x)
	// fmt.Printf("%T\n",&x)

	// y := &x

	// fmt.Println(*y,y)

	// *y = 32
	// fmt.Println(x)

	//---------------------------------


	x := 0
	foo(&x)
	fmt.Println(&x)
	fmt.Println(x)
	
}
func foo(y *int){
	fmt.Println(y)
	fmt.Println(*y)
	*y = 55
	fmt.Println(y)
	fmt.Println(*y)

}