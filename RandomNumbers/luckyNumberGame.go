package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurns    = 5
	noInput     = "Pick a number."
	nonInteger  = "Give an input is not a number."
	positiveNum = "Please pick a positive number."
	win         = "You Win!!!"
	lose        = "You lose.. Try Again ?"
)

func lucky() {

	rand.Seed(time.Now().UnixNano())

	if len(os.Args) != 2 {
		fmt.Println(noInput)
		return
	}

	guess, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(nonInteger)
		return
	}

	if guess < 0 {
		fmt.Println(positiveNum)
		return
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess + 1)
		if n == guess {
			fmt.Println(win)
			return
		}
	}
	fmt.Println(lose)
}
