package main

import (
	"fmt"
	"os"
	"strconv"
)

// run this program with a comman line argument
func main() {
	args := os.Args[1]

	celcius, _ := strconv.ParseFloat(args, 64)

	fahrenheit := (celcius * 1.8) + 32

	fmt.Printf("The %g celcius is converted to %f fahrenheit \n", celcius, fahrenheit)
	fmt.Printf("The %g celcius is converted to %g fahrenheit \n", celcius, fahrenheit)

}
