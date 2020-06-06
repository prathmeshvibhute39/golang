package main 

import "fmt"

type rushz struct{

	first string
	last string
	age int
}
type root struct{
	rushz
	first string
	ltk bool
}

func main(){
	p1 := root{
		rushz : rushz{
			first : "rushi",
			last :" vibhute",
			age : 20,
		},
		first: "guru",
		ltk : true,
	}


	fmt.Println(p1)
	fmt.Println(p1.first)
	fmt.Println(p1.rushz.first)



	//anpnymous composite literral

	p2 := struct{
		x string
		y string

	}{
		x: "fuck",
		y : "you",
	}
	fmt.Println(p2)

	
}