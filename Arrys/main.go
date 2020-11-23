package main

import "fmt"

const (
	winter = 1
	summer = 3
	yearly = winter + summer
)

func main() {
	//	moodly()
	//	arryaExample()
	arrayComapre()
	underlyingType()

}

func arryaExample() {
	var books [yearly]string
	books[0] = "kafka's revenge"
	books[1] = "Stay Golden"
	books[2] = "Everythingship"
	books[3] = books[0] + " 2nd Edition"

	// fmt.Printf("books :%T\n", books)
	// fmt.Println("books :", books)
	// fmt.Printf("books :%q\n", books)
	// fmt.Printf("books :%#v\n", books)

	//Other way of delcaring arrays
	names := [4]string{"Rida", "Ridha", "Inshira", "Sumaira"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	//Other way of delcaring arrays "Ellipsis Operator"
	names1 := [...]string{"Rida", "Sumaira"}
	for i := 0; i < len(names1); i++ {
		fmt.Println(names[i])
	}

}

func arrayComapre() {
	blue := [...]int{3, 6, 9}
	red := [3]int{3, 6, 9}
	fmt.Println("Are t hey same ?", blue == red)

	blue1 := [3]int{3, 6, 9}
	red1 := [3]int{3, 6}
	fmt.Println("Are t hey same ?", blue1 == red1)

	tem := blue1
	fmt.Println("Does it copied?", blue1 == tem)

}

func underlyingType() {

	type bookcase [3]int
	x := [3]int{1, 2, 3}
	blue := bookcase{1, 2, 3}
	fmt.Println("Are they equal", x == blue) // "TRUE" becaseu both the underlying types are same

	type cabinet [3]int
	red := cabinet{1, 2, 3}
	//	fmt.Println("Are they equal", blue == red) // comiler error, differetn types cannot be compared

	fmt.Println("	After Converting, are the equal ?", blue == bookcase(red)) // "TRUE" becasue we converted them

}
