package main

import (
	"fmt"
	"os"
)

type sample struct {
	name string
	test map[string]string
}

func main() {

	for _, env := range os.Environ() {
		fmt.Println(env)
	}

}
