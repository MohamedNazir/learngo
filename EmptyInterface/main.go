package main

import "fmt"

func main() {
	var any interface{}
	any = []int{2, 4, 6}
	any = map[string]int{"a": 1, "b": 2, "c": 3}
	any = "hello"
	any = 3 // --> here dynamic type of any is "int", dynamic value of any is "3"
	// 	any = any*2 ;

	any = any.(int) * 2 // need to extract it using type assertion.

	fmt.Println(any)

	nums := []int{3, 4, 5, 6}
	any = nums

	//   	_=len(any) // --> this won't compile

	_ = len(any.([]int))
}
