package main 

import ("fmt"
		"encoding/json"
		
 		
 		)


var jsonBlob = []byte(`[{"ID":1,"Name":"rushz","Colors":["Red","Blue","Black"]}]`)


type ColorGroup struct{
	ID int
	Name string
	Colors []string
}





func main(){

	var colors []ColorGroup

	err := json.Unmarshal(jsonBlob, &colors)

	if err != nil{
		fmt.Println("error: ",err)
	}

	fmt.Printf("%+v", colors)




}