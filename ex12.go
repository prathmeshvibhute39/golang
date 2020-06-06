package main

import "fmt"
 
type person struct{
	first string
	last string
	fav []string
}

func main(){
	p1 := person{
		first :"rushz",
		last :"root",
		fav : []string{
			"chocolate",
			"vannila",
			"strwaberry",
			"blueberry",
		},
	}
	p2 := person{
		first : "guru",
		last : "kahnde",
	}

	for k, v:= range p1.fav {
		fmt.Println(k,v)
	}

	fmt.Println(p2)


	m := map[string]person{
		p1.last :p1,
		p2.last :p2,
	}

	fmt.Println(m)


	

}