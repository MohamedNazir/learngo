package main

import "fmt"

// Duration type
type Duration int32

const (
	nanoseconds  Duration = 1
	microSeconds          = 1000 * nanoseconds
	milliseconds          = 1000 * nanoseconds
	seconds               = 1000 * nanoseconds
	minute                = 60 * seconds
	hours                 = 60 * minute
)

func main() {

	fmt.Println(nanoseconds, microSeconds, milliseconds, seconds, minute, hours)
}
