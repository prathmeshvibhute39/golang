package main

import "fmt"

const(

	_ = iota
	b = 1<< (iota *10)
	c = 1<< (iota *10)
	d = 1<< (iota *10)

)

func main(){


	// x := 4
	
	// fmt.Printf("%d\t\t%b\n",x,x)

	// y := x >> 1
	// fmt.Printf("%d\t\t%b\n",y,y)
	fmt.Printf("%d\t\t%b\n",b,b)
	fmt.Printf("%d\t\t%b\n",c,c)
	fmt.Printf("%d\t\t%b\n",d,d)
	
	
}