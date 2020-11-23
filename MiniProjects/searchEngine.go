package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	findWord()
	findWordWithLabledBreak()
	findWordWithLabledContinue()
}

const corpus = "lazy cat jumps again and again and again"

func findWord() {
	words := strings.Fields(corpus)
	query := os.Args[1:]
	for _, q := range query {
		for i, w := range words {
			if w == q {
				fmt.Printf(" # %d  --> %s \n", i+1, w)
				break
			}
		}
	}
}

// Labeled break
func findWordWithLabledBreak() {

	fmt.Println("-----------------------------------------------")
	words := strings.Fields(corpus)
	query := os.Args[1:]

queries:
	for _, q := range query {
		for i, w := range words {
			if w == q {
				fmt.Printf(" # %d  --> %s \n", i+1, w)
				break queries // this will break from the parent loop, without the label (queires), it will break only the current loop.
			}
		}
	}
}

// Labeled Continue
func findWordWithLabledContinue() {

	fmt.Println("-----------------------------------------------")
	words := strings.Fields(corpus)
	query := os.Args[1:]

queries:
	for _, q := range query {
		for i, w := range words {
			if w == q {
				fmt.Printf(" # %d  --> %s \n", i+1, w)
				continue queries
			}
		}
	}
}
