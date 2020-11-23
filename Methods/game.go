package main

import "fmt"

type game struct {
	title string
	price float64
}

func (g game) print() {
	fmt.Printf("{ %s : %f } \n", g.title, g.price)
}
