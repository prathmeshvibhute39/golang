package main

import "fmt"

type person struct{
	fname string
	lname string
}

type seAgent struct{
	person
	ltk bool
}

func (p person) speak(){
	fmt.Println(p.fname,p.lname)
}
func (s seAgent) speak(){
	fmt.Println(s.fname)
}


type human interface{
	speak()
}

func say(h human){
	h.speak()
}



func main(){


	p1 := person{"elon","musk"}
	s1 := seAgent{person{"rushz","vibhute"},true}

	say(p1)
	say(s1)
}