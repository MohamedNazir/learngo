package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()
	h := t.Hour()
	fmt.Printf("current hour is %d:\n", h)
	switch {
	case h > 0 && h < 12:
		fmt.Println("Good Morning")
	case h <= 12 && h <= 16:
		fmt.Println("Good Afternoon")
	case h > 16 && h <= 20:
		fmt.Println("Good Eveneing")
	default:
		fmt.Println("Good Night")
	}

}
