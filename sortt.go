package main

import ("fmt"
		"sort"
)

func main(){

	s := []int{4,8,5,6,9,7,5,2,6,8,5,3,5,8,96}

	t := []string{"f","d","g","k"}

	sort.Ints(s)
	sort.Strings(t)

	fmt.Println(s)
	fmt.Println(t)
}