package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {

	stringLiteralswithbackquote()
	stringConcatination()
	stringLength()

}

func stringLiteralswithbackquote() {
	s := "<html>\n\t<body>\"Hello\"</body>\n</html>"

	fmt.Println(s)

	s = `
	<html>
		<body>
			"Hello"
		</body>
	</html>
	`
	fmt.Println(s)
}

func stringConcatination() {
	name, last := "Carl", "sagan"
	name += " edward"
	fmt.Println(name + " " + last)

	eq := "Summation"
	sum := 3 + 5

	fmt.Println(eq + ": " + strconv.Itoa(sum))
}

func stringLength() {
	name := "rizan"
	fmt.Println("Length with only english charcters ", len(name))

	foreighname := "Önaßü"
	fmt.Println("Length with non-english charcters", len(foreighname)) // This len() function returns only the byte size, not the actual length

	// to fix this

	fmt.Println("Actual non-english length ", utf8.RuneCountInString(foreighname))
}
