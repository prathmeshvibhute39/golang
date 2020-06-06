package main

import "fmt"

func main(){
	s := rec(3)
	fmt.Println(s)
	fmt.Println(fac(3))
	fmt.Println(my())

}
//recursion
func rec(n int) int{
	if n==0{
		return 1
	}
	return n * rec(n-1)

}



//using for loop

func fac(n int) int{
	total:=1
	for; n > 0; n--{
		total*=n
	}
	return total
}




func my() int{
	x :=1
	
	for n :=4; n > 0; n--{
		x *= n
	}
	return x
}