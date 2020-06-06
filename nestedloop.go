package main

import "fmt"

func main(){


	for x:=0; x <=100; x++{

		if x % 3 == 0 && x % 5 ==0{
			fmt.Println("ping-pong")
		}else if x % 5 == 0{
			fmt.Println("pong")
		}else if x %3 ==0{
			fmt.Println("ping")
		}else{
			fmt.Println(x)
		}
	}


	for i:= 0; i <3 ; i++{

		for j:= 0; j<2 ;j++{
			for k:= 0; k < 1; k++{
				fmt.Println(i,j,k)
			} 
		}
	}
}