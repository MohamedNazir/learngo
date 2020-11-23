package main

import (
	"fmt"
	"strings"
)

type picasso struct {
	name string
	age  int
}

type author struct {
	name  string
	age   int
	email string
}

type song struct {
	title, artist string
}
type playList struct {
	genre string
	songs []song
}

func main() {

	var p picasso

	p.name = "Hel"
	fmt.Println(p.name)
	fmt.Printf("The type is %T, value is %v \n", p, p)

	author := author{name: "Rizan", age: 35, email: "test@testmail.com"}
	fmt.Printf("The type is %T, value is %v \n", author, author)

	person := Person{"asdf", 23, "asdfasd"}
	fmt.Printf("The type is %T, value is %v \n", person, person)

	song1 := song{"wonderwall", "oasis"}
	song2 := song{"super sonic", "oasis"}

	if song1.title == song2.title && song1.artist == song2.artist {
		fmt.Println("they are equal")
	} else {
		fmt.Println("Unequal objects")
	}

	// creating a playlist

	songs := []song{song1, song2}

	rockList := playList{genre: "rock", songs: songs}

	//clone := rockList

	// 	if (rockList == clone) // this comparing is not possible, because the struct used a slic which is not comparable.

	fmt.Printf("%-30s  %20s   %20s \n", "GENERE", "TITLE", "ARTIST")
	fmt.Println(strings.Repeat("-", 100))

	for _, s := range rockList.songs {
		fmt.Printf("%-30s  %20s   %20s \n", rockList.genre, s.title, s.artist)
	}

}
