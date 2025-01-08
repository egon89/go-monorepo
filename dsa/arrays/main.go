package main

import "fmt"

func main() {
	// Create a 1D array with 10 elements
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Print elements in reverse order
	fmt.Println("Array in reverse order:")
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}
