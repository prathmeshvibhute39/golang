package main

import "fmt"

func main(){

	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)
  	
  	go  send(eve,odd,quit)
  	receive(eve,odd,quit)
  	fmt.Println("about to quit")
}

func send(e,o chan<- int, q chan<- bool){

	for i:= 0; i<100 ;i++{
		if i %2 ==0{
			e <- i
		}else{
			o <- i
		}
	}
	close(q)

}


func receive(e,o <-chan int, q <-chan bool){
	for{
		select{

		case v := <-e:
			fmt.Println("even",v)
		case v := <-o:
			fmt.Println("odd",v)
	 	case v, ok := <-q:
	 		if !ok {
	 			fmt.Println("from ok",v)
	 		}else{
	 			fmt.Println("from ok",v)
	 		}

		}

	}
}