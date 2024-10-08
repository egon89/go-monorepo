package generics

import "fmt"

type MyNumber = float64

// constraint
type Number interface {
	int | float64
}

func GenericsExample() {
	fmt.Println("> genericsExample")
	m1 := map[string]int{"a": 100, "b": 200}
	fmt.Printf("sum: %v\n", sum(m1))
	fmt.Printf("sumNumberType: %v\n", sumNumber(m1))

	m2 := map[string]MyNumber{"a": 100.9, "b": 200.3}
	fmt.Printf("sumNumberType: %v\n", sum(m2))

	fmt.Printf("Compare 2 and 2: %v\n", compare(2, 2))
	fmt.Printf("Compare 2 and 1: %v\n", compare(2, 1))
	fmt.Printf("Compare 2 and 2.01: %v\n", compare(2, 2.01))
	fmt.Printf("Compare 2 and 2.0: %v\n", compare(2, 2.0))
}

func sum[T int | float64](m map[string]T) T {
	var total T
	for _, v := range m {
		total += v
	}

	return total
}

func sumNumber[T Number](m map[string]T) T {
	var total T
	for _, v := range m {
		total += v
	}

	return total
}

func compare[T comparable](a T, b T) bool {
	return a == b
}
