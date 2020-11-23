package main

import (
	"fmt"
	"os"
)

func main() {

	input := os.Args[1:]

	if len(input) < 1 {

		fmt.Println("Get me a key.....")
		return
	}
	key := input[0]

	dict := make(map[string]string)

	dict["h"] = "Hello"
	dict["a"] = "Apple"
	dict["b"] = "ball"
	dict["c"] = "Cat"
	dict["d"] = "Dog"
	dict["g"] = ""

	value, OK := dict[key]
	if !OK {
		fmt.Println("Key Not Found...")
		return
	}

	fmt.Printf("The valeu for %q is %q", key, value)
}
