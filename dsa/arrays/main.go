package main

import "fmt"

func main() {
	singleDimensionalArrayReverseOrder()
	multiDimensionalArray()
	sumAllElements()
	reverseASingleArray()
	reverseAMultiDimensionalArray()
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

func reverseAMultiDimensionalArray() {
	fmt.Println("> reverseAMultiDimensionalArray")
	matrix := [5][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
		{17, 18, 19, 20},
	}

	fmt.Println("matrix:", matrix)

	// reverse internal arrays
	for i := 0; i < len(matrix); i++ {
		for j, k := 0, len(matrix[i])-1; j < k; j, k = j+1, k-1 {
			// ik := matrix[i][k]
			// matrix[i][k] = matrix[i][j]
			// matrix[i][j] = ik
			fmt.Printf("(%d) changing matrix[%d][%d] from %d to %d and matrix[%d][%d] from %d to %d\n",
				i, i, j, matrix[i][j], matrix[i][k], i, k, matrix[i][k], matrix[i][j])
			matrix[i][k], matrix[i][j] = matrix[i][j], matrix[i][k]
		}
	}

	fmt.Println("reversed internal matrix:", matrix)

	// reverse root array
	for i, j := 0, len(matrix)-1; i < j; i, j = i+1, j-1 {
		fmt.Printf("changing matrix[%d] from %d to %d and matrix[%d] from %d to %d\n", i, matrix[i], matrix[j], j, matrix[j], matrix[i])
		matrix[i], matrix[j] = matrix[j], matrix[i]
	}

	fmt.Println("reversed matrix:", matrix)
	/*
		matrix: [[1 2 3 4] [5 6 7 8] [9 10 11 12] [13 14 15 16] [17 18 19 20]]
		(0) changing matrix[0][0] from 1 to 4 and matrix[0][3] from 4 to 1
		(0) changing matrix[0][1] from 2 to 3 and matrix[0][2] from 3 to 2
		(1) changing matrix[1][0] from 5 to 8 and matrix[1][3] from 8 to 5
		(1) changing matrix[1][1] from 6 to 7 and matrix[1][2] from 7 to 6
		(2) changing matrix[2][0] from 9 to 12 and matrix[2][3] from 12 to 9
		(2) changing matrix[2][1] from 10 to 11 and matrix[2][2] from 11 to 10
		(3) changing matrix[3][0] from 13 to 16 and matrix[3][3] from 16 to 13
		(3) changing matrix[3][1] from 14 to 15 and matrix[3][2] from 15 to 14
		(4) changing matrix[4][0] from 17 to 20 and matrix[4][3] from 20 to 17
		(4) changing matrix[4][1] from 18 to 19 and matrix[4][2] from 19 to 18
		reversed internal matrix: [[4 3 2 1] [8 7 6 5] [12 11 10 9] [16 15 14 13] [20 19 18 17]]
		changing matrix[0] from [4 3 2 1] to [20 19 18 17] and matrix[4] from [20 19 18 17] to [4 3 2 1]
		changing matrix[1] from [8 7 6 5] to [16 15 14 13] and matrix[3] from [16 15 14 13] to [8 7 6 5]
		reversed matrix: [[20 19 18 17] [16 15 14 13] [12 11 10 9] [8 7 6 5] [4 3 2 1]]
	*/
}
