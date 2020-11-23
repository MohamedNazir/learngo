package main

import (
	"fmt"
	"os"
)

//
//Args[0] -> is always the path of the program
//so need to take from args[1]

func main() {
	fmt.Printf("%#v\n", os.Args)
	firstArgument := os.Args[1]
	secondArgument := os.Args[2]

	fmt.Println("first Argument is :", firstArgument)
	fmt.Println("second Argument is :", secondArgument)

}
