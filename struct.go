package main 

import "fmt"

type rushz struct{

	first string
	last string
	age int
}

func main(){
	p1 := rushz{
		first :"rushz",
		last : "root",
		age : 20,
	}
	p2 := rushz{
		first :"guru",
		last : "kahnde",
		age : 55,
	}
	

	fmt.Println(p1)
	fmt.Println(p2)
}