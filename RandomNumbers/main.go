package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//	generateRandom() // this will generate the same reuslt even if you run 100 times.

	//generateRandomUsingSeed()
	lucky()
}

func generateRandom() {
	guess := 10
	for n := 0; n != guess; {
		n = rand.Intn(guess + 1)
		fmt.Printf("%d ", n) // the output is same, how many time we call it.
	}
	fmt.Println()
}

func generateRandomUsingSeed() {
	t := time.Now()
	rand.Seed(t.UnixNano())

	guess := 10
	for n := 0; n != guess; {
		n = rand.Intn(guess + 1)
		fmt.Printf("%d ", n) // the output is same, how many time we call it.
	}
	fmt.Println()
}
