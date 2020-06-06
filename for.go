package main

import "fmt"

func main(){


	for x := 0; x <= 100; x++{

		fmt.Println(x)
	}



	for {

		fmt.Println("infinite loop")
		break
	}
}