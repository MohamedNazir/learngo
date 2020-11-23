package main

import "fmt"

// RUN this program with "go run -race main.go "

func main() {

	go increment()
	go increment()
	fmt.Println(count)
}

var count int = 0

func increment() {

	if count == 0 {
		count = count + 1
	}
	fmt.Println(count)
}
