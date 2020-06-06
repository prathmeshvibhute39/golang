package main
import "fmt"
func main(){

	m := map[string]int{

		"rushz":9158393989,
		"guru": 156892585,
		"rahul" : 454554545,
	}

	fmt.Println(m["rushz"])

	v ,ok := m["akshay"]
	fmt.Println(v)
	fmt.Println(ok)

	if v ,ok := m["guru"]; ok{
		fmt.Println(v)
	}


	m["ak"] = 454554

	fmt.Println(m)


	for k,v := range m{

		fmt.Println(k,v)
	}

	delete (m,"ak")
	fmt.Println(m)

}