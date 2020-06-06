package main

import "fmt"

func main(){


	x := []string{"rushi","vibhute","vanila"}
	y := []string{"komuuu","vibhute","vanila"}

	fmt.Println(x)
	fmt.Println(y)

	xy := [][]string {x,y}
	fmt.Println(xy)

	for k, v := range x{
		fmt.Println(k,v)
	}

}	

