package main

import "fmt"

type list []*game

func (l list) print() {
	for _, it := range l {
		it.print()
	}
}

func (g game) print() {
	fmt.Printf("{ %s : %f } \n", g.title, g.price)
}
