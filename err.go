package main

import (
	"fmt"
	"os"
	"io/ioutil"
	//"strings"
)

func main(){

	//create and edit file
	f, err := os.Create("names.txt")
		if  err != nil{
			fmt.Println("error: ", err)
			return
		}
		defer f.Close()
		
		r := strings.NewReader("james bond")
		io.Copy(f, r)

	//------------------read from file

	// f, err := os.Open("names.txt")
	// 	if err != nil {
	// 		fmt.Println("error: ",err)
	// 	}
	// 	defer f.Close()

	// bs ,err := ioutil.ReadAll(f)
	// 	if err != nil {
	// 		fmt.Println("error: ",err)
	// 	}

	// fmt.Println(string(bs))


}