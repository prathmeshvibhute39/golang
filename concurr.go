package main

import (
	"fmt"
	"runtime"
	"sync"
	//"os"
)

var wg sync.WaitGroup

func main(){

	fmt.Println("os\t",runtime.GOOS)
	fmt.Println("arch\t", runtime.GOARCH)
	fmt.Println("cpu\t",runtime.NumCPU())
	fmt.Println("gorutine\t", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	bar()

	fmt.Println("cpu\t",runtime.NumCPU())
	fmt.Println("gorutine\t", runtime.NumGoroutine())
	wg.Wait()
}



func foo(){

	for i:=0; i<=5; i++{
		fmt.Println(i)
	}
	wg.Done()
}
func bar(){

	for i:=5; i<=10; i++{
		fmt.Println(i)
	}
}