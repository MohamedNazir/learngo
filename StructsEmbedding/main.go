package main

import "fmt"

//  INHERITANCE
//  is a relationship  A book is a text An Article is a text

//  Go DOES NOT use is a relationshop, It uses HAS a Relationship (COMPOSITION)

type text struct {
	title string
	words int
}

type book struct {
	text
	isbn string
}

func main() {

	normalWay()
	anonnumusWay()

}

func normalWay() {

	type book struct {
		text text
		isbn string
	}
	moby := book{
		text: text{title: "The Moby dick", words: 15423},
		isbn: "as43-3e45-asdf-2342",
	}

	fmt.Printf(" Tile : %s, Words: %d, ISNB: %s \n", moby.text.title, moby.text.words, moby.isbn)
}

func anonnumusWay() {

	// Annonymus fields  -> A struct field without a name is called annonymus field.
	// from line no 58, if we remove the type "text" and run the program, it will still work,
	// because go will find its type from its name if not defined.
	// in this case we remove the "text" from the printf function also, and ca directly access like moby.title, moby.words

	type book struct {
		text
		isbn string
	}
	moby := book{
		text: text{title: "The Moby dick", words: 15423},
		isbn: "as43-3e45-asdf-2342",
	}

	fmt.Printf(" Tile : %s, Words: %d, ISNB: %s \n", moby.title, moby.words, moby.isbn)
}
