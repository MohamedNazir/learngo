package main

import "fmt"

func main() {

	matrix := [2][3]int{
		[3]int{1, 2, 3},
		[3]int{4, 5, 6},
	}
	fmt.Println(matrix)

	matrix1 := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(matrix1)

	matrix2 := [...][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(matrix2)

	calculation()
}

func calculation() {
	students := [...][3]float64{
		{5, 6, 8},
		{9, 3, 2},
	}

	var sum float64
	for _, stud := range students {
		for _, marks := range stud {
			sum += marks
		}
	}

	fmt.Printf("The Avg mark is %f \n", sum)

}
