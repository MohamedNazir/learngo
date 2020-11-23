package main

import (
	"fmt"
	"reflect"
)

func main1() {
	var speed int
	speed = 120

	fmt.Println("Integer Variables")
	fmt.Println(speed)
	fmt.Printf("The value is %d and type is %T\n", speed, speed)
	fmt.Println(reflect.TypeOf(speed))

	var mileage int = 54
	fmt.Printf("The value is %d and type is %T\n", mileage, mileage)

	var mileage32 int32 = -20
	fmt.Printf("The value is %d and type is %T\n", mileage32, mileage32)

	var mileage64 int64 = 120
	fmt.Printf("The value is %d and type is %T\n", mileage64, mileage64)

	var usi int16 = -32
	fmt.Printf(" value type is %T\n ", usi)

	var (
		currentspeed = 100
		prevspeed    = 50
	)

	currentspeed, prevspeed = prevspeed, currentspeed

	fmt.Println(currentspeed, prevspeed)

}
