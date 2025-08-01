package main

// slice
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

// variadic function: accepts a variable number of arguments
func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		// append: returns a new slice with all the items in it.
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
			continue
		}

		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}
