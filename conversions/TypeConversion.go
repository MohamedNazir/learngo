package main

import (
	"fmt"
)

func main() {
	fmt.Println("TYPE CONVERSIONS")

	speed := 100
	force := 2.5
	speed = speed * int(force)
	fmt.Println(speed) // prints a wrong answer , because converting float to int leads to loss of data

	speed1 := 100
	force1 := 2.5
	var result float64 = float64(speed1) * force1
	fmt.Println(result)

	var apple int
	var orange int32
	//apple = orange // gives error

	apple = int(orange)
	fmt.Println(apple)

	orange = 65
	color := string(orange)

	fmt.Println(color) // this will print a string "A" becaue ist was the ascii value;
}
