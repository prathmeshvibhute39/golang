package main

import "fmt"

const usmax float64 = 65535
type car struct{
	gas_padel uint16
	break_pedal uint16
	steering_wheel int16
	top_speed float64
}

func (c *car) kmh() float64{
	c.top_speed = 500
	return float64(c.gas_padel) * (c.top_speed/usmax)
}

func (c *car) new_top_spedd(newspeed float64){
	c.top_speed = newspeed
}
func main(){

	a_car := car{
		gas_padel : 22341,
		break_pedal :0,
		steering_wheel :12564,
		top_speed: 225}
	fmt.Println(a_car.gas_padel)
	fmt.Println(a_car.kmh())
	a_car.new_top_spedd(500)
	fmt.Println(a_car.kmh())


}
