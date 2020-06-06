package main

import "fmt"

func main(){

	x:= [5]int{1,2,3,4,5}

	fmt.Println(x)
	fmt.Printf("%T\n",x)

	for k, v := range x{
		fmt.Println(k,v)
	}


	y := []int{1,2,3,4,5,6,7,8,9,10}

	for i, j := range y{
		fmt.Println(i,j)
	}


	a := y[:5]
	b := y[5:]
	fmt.Println(a)
	fmt.Println(b)


	k := []int{42,43,44,45,46,47,48,49,50,51}
	k = append(k,52)
	fmt.Println(k)
	k = append(k,53,54,55)
	fmt.Println(k)

	l := []int{56,57,58,59,60}
	k = append(k,l...)
	fmt.Println(k)

	z := append(k[:3],k[6:]...)
	fmt.Println(z)



	state := [] string {"root","rushz","guru","rahul"}
	fmt.Println("length of sclice: ",len(state))
	fmt.Println("capacity of slice: ",cap(state))


	for p, q := range state{
		fmt.Println(p,q)
	}


	n := []string{"james","Bond","shaken not striied"}
	m  := []string{"miss","root","rushz"}

	for c:=3 ; len(n) >=c ;c++{
		for d:=3 ; len(m) >=d ;d++{
			fmt.Println(n,m)
		}

	}

	xx := [][]string {n,m}
	fmt.Println(xx)


	xxx := map[string]string{

		"first" : "Bond_james",
		"second" : "IceCream",
	}

	fmt.Println(xxx)

	for g,h := range xxx{
		fmt.Println(g,h)
	}

	delete(xxx,"first")
	fmt.Println(xxx)
}