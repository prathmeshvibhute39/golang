package main 

import ("fmt"
		"encoding/json"
		"os"
 		
 		)

type ColorGroup struct{
	ID int
	Name string
	Colors []string
}

func main(){
	group := ColorGroup{1,"rushz",[]string{"Red","Blue","Black"}}

	b , err := json.Marshal(group)
	if err != nil {
		fmt.Println("error: ",err)
	}
	os.Stdout.Write(b)

}