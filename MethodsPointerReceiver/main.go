package main

import "fmt"

func main() {

	tennis := game{"Tennis", 20.99}

	tennis.discount(.5)
	tennis.print()

	// you also write it like this,(I mean passing the & operator) but it was not required go does it internally
	carrom := game{"CARROM", 10.0}
	(&carrom).discount(.5)
	carrom.print()
}

type game struct {
	name  string
	price float64
}

// why we are using pointer here is just for consistany, though we are not modifuying it.
func (g *game) print() {
	fmt.Printf("%-15s : $%2f\n", g.name, g.price)
}

// pass the game as pointer
func (g *game) discount(ratio float64) {
	g.price *= (1 - ratio)
}
