package main

import "fmt"

type gram float64
type ounce float64

type (
	grm int
	kgm grm
	ton kgm
)

func main() {

	var g gram = 10
	var o ounce
	o = ounce(g) * 0.035274
	fmt.Printf("%g grams is %g ounce\n ", g, o)

	var salt grm = 10
	var apple kgm = 20
	var truck ton = 10

	fmt.Printf("salt : %d, apple :%d truck : %d \n", salt, apple, truck)

}
