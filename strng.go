package main

import "fmt"

func main(){

	skills := "Hello ,rushz"
	fmt.Println(skills)
	// fmt.Printf("%T\n",s)
	// x := []byte(s)
	// fmt.Println(x)
	// fmt.Printf("%T\n",x)

	fmt.Print("y/n : " )
	var ans string
	fmt.Scanln(&ans)
	//fmt.Println("you entered: ",ans)


	if ans == "y"{
		fmt.Println("good boy")
	}else{
		fmt.Println("fuck you")
	}
}