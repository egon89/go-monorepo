package main

import "fmt"

func main() {
	singleDimensionalArrayReverseOrder()
	multiDimensionalArray()
	sumAllElements()
	reverseASingleArray()
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

func sumAllElements() {
	fmt.Println("> sum all elements in an array:")

	numbers := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sum := 0
	for _, value := range numbers {
		sum += value
	}

	fmt.Println("sum of elements:", sum)
}

func reverseASingleArray() {
	fmt.Println("> reverse a single array:")

	numbers := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
		/*
			same as:
			iValue := numbers[i]
			numbers[i] = numbers[j]
			numbers[j] = iValue
		*/
	}

	fmt.Println("reversed array:", numbers)
}
