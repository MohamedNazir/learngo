package main

import "fmt"

type book struct {
	title string
	price float64
}

func printBook(b book) {
	fmt.Printf("{ %s : %f } \n", b.title, b.price)
}
