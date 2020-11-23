package main

import (
	"fmt"
	"sort"
)

func main() {
	//basic()
	//	sliceExpressions()
	pagination()
}

func basic() {
	games := []string{"pokeman", "sims"}
	greet := []string{"hello", "doctor"}

	games = append(games, "Mytest")
	games = append(games, "Another", "element", "added")

	games = append(games, greet...)

	fmt.Println(games)

	numbers := []int{2, 4, 8, 1, 3, 0, 9, 12, 51, 23}

	fmt.Println(numbers)

	sort.Ints(numbers)
	fmt.Println(numbers)
}

func sliceExpressions() {

	names := []int{1, 2, 3, 4, 5, 6, 7, 0, 4, 22, 15, 27, 23}
	//msg := []byte{'h', 'e', 'l', 'l', 'o'}

	let := names[0:4]
	tee := append(let, 777)
	fmt.Println(let)
	fmt.Println(tee)

	fmt.Println(names[:])

}
