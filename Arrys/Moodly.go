package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func moodly() {

	if len(os.Args[1:]) != 1 {
		fmt.Println("[Your name ?]")
		return
	}
	name := os.Args[1]
	moods := [...]string{"Happy :)", "Good ", "Awesome", "Sad", "Terrible", "Bad"}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(moods))

	fmt.Printf("%s feeling %q now.\n", name, moods[n])
}
