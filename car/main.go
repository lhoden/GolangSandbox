package main

import (
	"fmt"
)

// We need to use *car because it's a pointer receiver (so we're referring to an object that already exists)
func (c car) kmh() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh / usixteenbitmax)
}

func (c *car) mph() float64 {
	c.topSpeedKmh = 500
	return float64(c.gasPedal) * (c.topSpeedKmh / usixteenbitmax / kmhMultiple)
}

func (c *car) newTopSpeed(newSpeed float64) {
	c.topSpeedKmh = newSpeed
}

func main() {
	//Creating an object
	aCar := car{gasPedal: 22341,
		brakePedal:    0,
		steeringWheel: 12561,
		topSpeedKmh:   225.0}

	fmt.Println(aCar.gasPedal)
	fmt.Println(aCar.kmh())
	fmt.Println(aCar.mph())
	aCar.newTopSpeed(500)
	fmt.Println(aCar.kmh())
	fmt.Println(aCar.mph())
}

type car struct {
	gasPedal      uint16 // unsigned integer means min 0 max 65535
	brakePedal    uint16
	steeringWheel int16
	topSpeedKmh   float64
}

const usixteenbitmax float64 = 65535
const kmhMultiple float64 = 1.60934
