package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {

	go sortInts()
	go sortStrings()
	go sortFloats()
	go sortStruct()
	go sortListByTitle()

	time.Sleep(1000 * time.Millisecond)
}

func sortInts() {
	s := []int{4, 2, 3, 1}
	sort.Ints(s)
	fmt.Println(s) // [1 2 3 4]
}

func sortStrings() {
	s := []string{"zebra", "cat", "ink", "fish"}
	sort.Strings(s)
	fmt.Println(s) // [cat fish ink zebra]
}

func sortFloats() {
	s := []float64{23.63, 03.23, 234.234, 35.126}
	sort.Float64s(s)
	fmt.Println(s) // [3.23 23.63 35.126 234.234]
}

type person struct {
	name string
	age  int
	city string
}

func sortStruct() {

	family := []struct {
		Name string
		Age  int
	}{
		{"Alice", 23},
		{"David", 2},
		{"Eve", 2},
		{"Bob", 25},
	}

	// Sort by age, keeping original order or equal elements.
	sort.SliceStable(family, func(i, j int) bool {
		return family[i].Age < family[j].Age
	})
	fmt.Println(family) //[{David 2} {Eve 2} {Alice 23} {Bob 25}]
}

func sortListByTitle() {
	p1 := person{"abd", 32, "london"}
	p2 := person{"zid", 20, "paris"}
	p3 := person{"Sam", 45, "luxembourg"}
	p4 := person{"Riz", 10, "amsterdam"}

	l := []person{p1, p2, p3, p4}

	sort.Sort(ByAge(l))
	fmt.Println(l) //[{Riz 10 amsterdam} {zid 20 paris} {abd 32 london} {Sam 45 luxembourg}]
}

// ByAge implements sort.Interface based on the Age field.
type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
