package main

import (
	"fmt"
	"path"
)

const seperator string = "/"

func main() {

	dir, filename := path.Split("test/src/css/main.css")
	fmt.Println(dir, filename)

	_, filename1 := path.Split("css/main.css")
	fmt.Println("file Name is :", filename1)

	dir1, _ := path.Split("js/main.css")
	fmt.Println("directory Name is :", dir1)
}
