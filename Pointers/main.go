package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("..... ARRAYS")
	arrays()
	slices()
	maps()
	structEx()

}

func arrays() {

	nums := [...]int{1, 2, 3, 4}
	fmt.Printf("The address is %p\n", &nums)
	incr(nums)
	fmt.Println(nums)

	incrByPtr(&nums)
	fmt.Println(nums)
}

func slices() {
	names := []string{"up", "down", "left", "right"}
	fmt.Println("Before update", names)
	convertToUpper(names)
	fmt.Println("after update without pointers", names)
}

func maps() {
	confused := map[string]int{"one": 2, "two": 3}

	fmt.Println(confused)
	fix(confused)
	fmt.Println(confused)
}

type house struct {
	name string
	room int
}

func structEx() {
	myHouse := house{"Rizan", 4}

	addRoom(myHouse)
	fmt.Println("House without Pointer ", myHouse)

	addRoomPtr(&myHouse)
	fmt.Println("House with  Pointer ", myHouse)

}

func incr(nums [4]int) {
	fmt.Printf("The address is %p\n", &nums)
	for i := range nums {
		nums[i]++
	}

}

func incrByPtr(nums *[4]int) {
	fmt.Printf("The address is %p\n", &nums)
	for i := range nums {
		nums[i]++
	}

}

//No need to use pointers for map and slice
func convertToUpper(names []string) {
	for i := range names {
		names[i] = strings.ToUpper(names[i])
	}
}

//No need to use pointers for map and slice
func fix(m map[string]int) {

	m["one"] = 1
	m["two"] = 2

}

func addRoom(h house) {
	h.room = h.room + 1
}

func addRoomPtr(h *house) {
	h.room = h.room + 1
}
