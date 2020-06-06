package main

import ("fmt"
		"encoding/json"
		
)


var data = []byte(`[{"ID":1,"Name":"rushz","Color":["Red","Blue","Black"]},
				 {"ID":2,"Name":"guru","Color":["Red","Blue","Black"]},
			   	{"ID":3,"Name":"rahul","Color":["Red","Blue","Black"]},
				{"ID":4,"Name":"apk","Color":["Red","Blue","Black"]}]`)

type User struct{
	ID int
	Name string
	Color []string
}

func main(){
	var d []User

	err := json.Unmarshal(data,&d)
	if err != nil{
		fmt.Println("error:",err)
	}
	fmt.Printf("%+v", d)    //%v	the value in a default format
							//when printing structs, the plus flag (%+v) adds field names


}