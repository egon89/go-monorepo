package closures

import "fmt"

func ClosureExample() {
	fmt.Println("> closureExample")
	total := func() int {
		return 2 + 5
	}()
	fmt.Printf("total as value: %d\n", total)

	sum := func(a, b int) int {
		return a + b
	}
	fmt.Printf("sum as function: %d\n", sum(2, 6))

	returnValueFromClosure := returnAFunction()
	fmt.Printf("returned value from a closure: %d\n", returnValueFromClosure())

	strategy := strategyFunction("apple")
	fmt.Printf("%s\n", strategy("."))

	strategy2 := strategyFunction("banana")
	fmt.Printf("%s\n", strategy2("!"))

	/*
		total as value: 7
		sum as function: 8
		returned value from a closure: 2
		length less than 6: apple.
		length greater than or equal to 6: banana!
	*/
}

func returnAFunction() func() int {
	return func() int {
		return 2
	}
}

func strategyFunction(word string) func(suffix string) string {
	if len(word) < 6 {
		return func(suffix string) string {
			return fmt.Sprintf("length less than 6: %s%s", word, suffix)
		}
	}

	return func(suffix string) string {
		return fmt.Sprintf("length greater than or equal to 6: %s%s", word, suffix)
	}
}
