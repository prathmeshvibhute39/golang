package main

import "fmt"
func main(){
	foo()
	bar("rushz")
	
	
	//functionn expression

	f := func(){

		fmt.Println("function expression")
	}
	f()

	fmt.Println(re())
	fmt.Println(hello())


}
func foo(){
	fmt.Println("i am from foo")
}

func bar(s string){
	fmt.Println("hello",s)
}


func re() int{
	return 451
}

func fu() string{
	return "whats going on"
}


func hello() string {

	return "this is inner one " +fu()
	
		
	
}





