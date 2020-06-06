package main

import ("fmt"
		"math"
)

type square struct{
	height float64
	width float64
}

type circle struct{
	radius float64
}

func (s square) area() float64 {
	return s.height * s.width

}

func (c circle) area() float64{
	return math.Pi * c.radius* c.radius

}

type shape interface{
	area() float64
}

func info(s shape){
	fmt.Println(s.area())
}

func main(){
	

	s := square{12.14,45.232}
	// ss :=s.area()
	// fmt.Println(ss)
	info(s)

	c := circle{12.244}
	// cc := c.area()
	// fmt.Println(cc)
	info(c)




	//anonymous func
	func() {
	  fmt.Println("Anonymous ")
	}()


	//assigning the func to variable

	f := func() {
		fmt.Println("This one assigned to variable")
	}
	f()


	//func return fucn


	func re() func() int{
		return func()int {
			return 45
		}
	}

	 f := re()
	 fmt.Println(f())

}