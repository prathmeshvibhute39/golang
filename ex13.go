package main

import "fmt"

type vehicle struct{
	doors int
	color string
}
type truck struct{
	vehicle
	fourWheel bool
}
type seden struct{
	vehicle
	luxury bool
}


func main(){

	x := truck{
		vehicle : vehicle{
			doors : 4,
			color :" red",
		},
		fourWheel : true,
	} 

 	y := seden{
 		vehicle : vehicle{
 			doors: 4,
 			color : "black",
 		},
 		luxury : true,
 	}

 	fmt.Println(x)
 	fmt.Println(y)
 	fmt.Println(x.doors)
 	fmt.Println(y.color)


 	//ananomys struct

 	xx := struct{

 		name string
 		color string
 		num int
 	}{
 		name :"rushz",
 		color  :"black",
 		num :3,
 	}
 	fmt.Println(xx)
 	m := map[string]struct{
 		xx.name : xx,
 	}
 	for i, v := range m{
 		fmt.Println(i,v)
 	}

}