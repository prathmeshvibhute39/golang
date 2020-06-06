package main

import ("fmt"
		"net/http"
)
func main(){
	resp, err := http.Get("https://www.google.com")

	fmt.Println(resp)	
	if err != nil{
		fmt.Println(err)
	}else{
		defer resp.Body.Close()
		fmt.Println("Conenction closed")
	}

}