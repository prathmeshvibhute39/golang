package main
 
import "fmt"

func main(){

	for i:= 10; i <= 100;i++{

		fmt.Printf("%v module %v\n",i,i%4)
	}
}