package main

import "fmt"

func main() {
	singleDimensionalArrayReverseOrder()
	multiDimensionalArray()
}

func singleDimensionalArrayReverseOrder() {
	fmt.Println("> array in reverse order:")
	// Create a 1D array with 10 elements
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Print elements in reverse order
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Printf("%d ", arr[i])
	}
}

func multiDimensionalArray() {
	fmt.Println("> traversing the multi-dimensional array")

	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println("using for loop")
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("element at [%d][%d]: %d\n", i, j, matrix[i][j])
		}
	}

	fmt.Println("using range")
	for i, second := range matrix {
		for j, value := range second {
			fmt.Printf("element at [%d][%d]: %d\n", i, j, value)
		}
	}
}
