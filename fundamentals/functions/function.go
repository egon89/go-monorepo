package functions

import (
	"errors"
	"fmt"
)

func FunctionExample() {
	println("> functionExample")
	result, err := sum(40, 15)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("sum result: %d\n", result)

	fmt.Printf("sum of numbers: %d\n", sumNumbers(7, 7, 7, 7, 7, 7, 7))

	numbersToSum := []int{7, 7, 7}
	fmt.Printf("sum of numbers from slice: %d\n", sumNumbers(numbersToSum...))

	/*
		the sum is greater than 50
		sum result: 0
		sum of numbers: 49
		sum of numbers from slice: 21
	*/
}

func sum(a, b int) (int, error) {
	sum := a + b
	if sum > 50 {
		return 0, errors.New("the sum is greater than 50")
	}

	return sum, nil
}

// variadic function
func sumNumbers(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}
