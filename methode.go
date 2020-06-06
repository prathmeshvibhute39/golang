package main

import "fmt"

type person struct{
	first string
	last string
}
type secretAgent struct{
	person
	ltk bool
}


func main(){
	sa1 := secretAgent{
		person:person{
			"james",
			"Bond",
		},
		ltk : true,
	}
	sa2 := secretAgent{
		person:person{
			"jamy",
			"Bondy",
		},
		ltk : true,
	}


	//fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
}
     //receiver
func (s secretAgent) speak(){
	fmt.Println("I am",s.first)
}

