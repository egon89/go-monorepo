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

	/*
		the sum is greater than 50
		sum result: 0
	*/
}

func sum(a, b int) (int, error) {
	sum := a + b
	if sum > 50 {
		return 0, errors.New("the sum is greater than 50")
	}

	return sum, nil
}
