package main

import ("fmt"
		"encoding/json"
		"os"
)

type User struct{
	ID int
	Name string
	Color []string
}

func main(){
		u1 := User{1,"rushz",[]string{"Red","Blue","Black"}}
		u2 := User{2,"guru",[]string{"Red","Blue","Black"}}
		u3 := User{3,"rahul",[]string{"Red","Blue","Black"}}
		u4 := User{4,"apk",[]string{"Red","Blue","Black"}}

		//var u User
		data := []User{u1,u2,u3,u4}
		d, err := json.Marshal(data)

		if err != nil{
			fmt.Println("error: ",err)
		}

		os.Stdout.Write(d)




}