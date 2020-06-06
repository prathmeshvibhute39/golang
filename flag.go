package main

import (
	"fmt"
	"flag"
)

func main(){
	minuso := flag.Bool("-o", false, "o")
	minusc := flag.Bool("c", false, "c")
	minusk := flag.Int("k", 0, "an int")

	fmt.Println("-o", *minuso)
	fmt.Println("-c", *minusc)
	fmt.Println("-k", *minusk)

	for index, val := range flag.Args(){
		fmt.Println(index, ":", val)
	}

}