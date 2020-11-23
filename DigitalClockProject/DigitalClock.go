package main

import (
	"fmt"
)

func main() {

	digits := [...]Placeholder{
		Zero, One, Two, Three, Four, Five, Six, Seven, Eight, Nine,
	}

	fmt.Println(digits)

	// for line := range digits[0] {
	// 	// Print a line for each placeholder in digits
	// 	for digit := range digits {
	// 		fmt.Print(digits[digit][line], "  ")
	// 	}
	// 	fmt.Println()
	// }
	// for i = 0; i < 1; i++ {
	// 	screen.Clear()
	// 	hour, minute, sec := time.Now().Clock()
	// 	fmt.Printf("%d : %d : %d \n", hour, minute, sec)
	// }
}
