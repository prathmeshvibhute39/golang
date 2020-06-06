package main

import "fmt"

func main(){
	func(){
		fmt.Println("I am anonymouse")
	}()

	func(x int){
		fmt.Println("I am anonymouse",x)
	}(42)

	

}

