package main

import "fmt"

func main(){

	s:= []int{1,2,3,4,5,6,7,8,9,10}
	t:= []int{1,2,3,4,5,6,7,8,9,10,11}

	defer fmt.Println(foo(s...))
	fmt.Println(bar(t))
}

func foo(n...int) int{
	total := 0
	for _,v := range n{

		total += v
		
	}
	return total
}
//-----------------------------------------
func bar(n []int)int{
	tot := 0
	for _, v := range n{
		tot += v
	}
	return tot
}