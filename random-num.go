package main

import (
	"fmt"
	"math/rand"
)

func main(){
	min := 1
	max := 100

	sec := rand.Intn(max-min) + min

	fmt.Println(sec)

}