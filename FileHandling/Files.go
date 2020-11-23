package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("Enter the directory")
		return
	}

	files, err := ioutil.ReadDir(args[0])
	if err != nil {
		fmt.Print(err)
	}
	for _, f := range files {
		fmt.Println(f.Name())
		fmt.Println(f.Size())
		fmt.Print(f.ModTime())
	}

}
