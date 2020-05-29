package main

import "fmt"

const usmax float64 = 65535
type car struct{
	gas_padel uint16
	break_pedal uint16
	steering_wheel int16
	top_speed float64
}

func (c car) kmh() float64{
	return float64(c.gas_padel) * (c.top_speed/usmax)
}

func main(){

	a_car := car{
		gas_padel : 22341,
		break_pedal :0,
		steering_wheel :12564,
		top_speed: 225}
	fmt.Println(a_car.kmh())

}
