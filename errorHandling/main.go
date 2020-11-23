package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := strconv.Itoa(65)
	fmt.Println(s)

	n, error := strconv.Atoi("42")
	fmt.Println(n)
	fmt.Println(error)

	n1, err := strconv.Atoi("4a")
	if err != nil {
		fmt.Println("ERROR :", err)
		return
	}
	fmt.Println("SUCCESS: Converted to", n1)

}
