package main

import "fmt"

type person struct{
	first string
	last string
	age int
}

func (p person )speak(){
	fmt.Println("I am,",p.first )
}

func main(){
	s := person{
		"rushz",
		"root",
		45,
	}

	s.speak()

}