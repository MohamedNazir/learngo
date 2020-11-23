package main

import "fmt"

func main() {

	score := 5

	if score > 3 {
		fmt.Println("good")
	} else {
		fmt.Println("Bad")
	}

	fizzBuzz()
}

func fizzBuzz() {

	x := 15

	if x%3 == 0 && x%5 == 0 {
		fmt.Println("Fizz Buzz")
		return
	} else if x%3 == 0 {
		fmt.Println("Fizz")
		return
	} else if x%5 == 0 {
		fmt.Println("Buzz")
		return
	} else {
		fmt.Println(" Not of any")
		return
	}
}
