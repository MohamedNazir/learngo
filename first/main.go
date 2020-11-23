package main

import (
	"fmt"
	f "fmt"
	"time"
)

const authorName = "Mohamed Nazir"

var hello = "HELLO"

func main() {
	fmt.Println("Hello, Gopher")
	fmt.Println(printAuthorName())
	printBye()
	fmt.Println(time.Now().Clock())
	f.Println("Duplicate Importing using alias")

	if 5 > 1 {
		f.Println(" 5 is bigger than 1")
	}
}

func printAuthorName() string {
	return authorName
}
