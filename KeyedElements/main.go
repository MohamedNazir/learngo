package main

import "fmt"

func main() {
	concepts()
	exampleProgram()
}

func concepts() {
	reates := [3]float64{
		2: 3.2,
		0: 0.5,
		1: 0.3,
	}
	fmt.Println(reates)

	reates1 := [...]float64{
		1: 0.3,
		4: 2.4,
	}
	fmt.Println(reates1)

	reates2 := [...]float64{
		5: 0.3,
		3.143,
		0: 2.4,
	}
	fmt.Println(reates2)
}

func exampleProgram() {

	const (
		ETH = iota
		WAN
	)

	rates := [...]float64{
		ETH: 25.5,
		WAN: 120.5,
	}

	fmt.Printf(" 1 BTC is %g ETH \n", rates[ETH])
	fmt.Printf(" 1 BTC is %g WAN \n", rates[WAN])
}
