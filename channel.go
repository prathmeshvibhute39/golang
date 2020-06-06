package main

import(
	"fmt"
	//"sync"
	"runtime"
	//"time"
)

// func do(x int) int{
// 	return x*5
// }
func main(){

	// ch := make(chan int)
	// go func(){
	// 	ch <- do(5)
	// }()

	// fmt.Println(<-ch)



	//---------------------------


	// c := make(chan int)
	// go func(){
	// 	for i:=0; i <10; i++{
	// 		c <- i
	// 	}
	// }()

	// go func(){
	// 	for{

	// 		fmt.Println(<-c)
	// 		}

	// 	}()
	// 	time.Sleep(time.Second*2)


	//---------------


	// c := make(chan int)
	// fmt.Println("goroutine",runtime.NumGoroutine())

	// go func(){
	// 	c <- 42
	// }()

	// fmt.Println("goroutine",runtime.NumGoroutine())

	// fmt.Println(<-c)


	//------------buffer value

	// c := make(chan int, 1)
	// fmt.Println("goroutine",runtime.NumGoroutine())

	
	// 	c <- 42
	

 //    fmt.Println(<-c)



 //--------------------unsuccessful buffer


 // c := make(chan int, 1)
	// fmt.Println("goroutine",runtime.NumGoroutine())

	
	// 	c <- 42
	// 	c <- 43
	

 //    fmt.Println(<-c)

 //----------------successfu; buffer

 // c := make(chan int, 2)
	// fmt.Println("goroutine",runtime.NumGoroutine())

	
	// 	c <- 42
	// 	c <- 43
	

 //    fmt.Println(<-c)
 //    fmt.Println(<-c)

//-------------------send only channel

 // c := make(chan <-int, 2)
	// fmt.Println("goroutine",runtime.NumGoroutine())

	
	// 	c <- 42
	// 	c <- 43
	

 //    fmt.Println(<-c)
 //    fmt.Println(<-c)

 //-------------------receive only channel

 c := make(<-chan int, 2)
	fmt.Println("goroutine",runtime.NumGoroutine())

	
		c <- 42
		c <- 43
	

    fmt.Println(<-c)
    fmt.Println(<-c)



}