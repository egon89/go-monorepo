package main

import "fmt"

const fund = "> Fundamentals"

type ID int

// global variables
var (
	b bool    = true
	c int     // default value: 0
	d string  // default value: ""
	e float64 = 1.2
	f ID      = 1
)

func main() {
	fmt.Println(fund)
	printGlobalVariables()
	arrayExample()
	sliceExample()
}

func printGlobalVariables() {
	fmt.Printf("b type is %T and value %v\n", b, b)
	fmt.Printf("c type is %T and value %v\n", c, c)
	fmt.Printf("d type is %T and value %v\n", d, d)
	fmt.Printf("e type is %T and value %v\n", e, e)
	fmt.Printf("f type is %T and value %v\n", f, f)
	/*
	  b type is bool and value true
	  c type is int and value 0
	  d type is string and value
	  e type is float64 and value 1.2
	  f type is main.ID and value 1
	*/
}

func arrayExample() {
	var myArray [3]int
	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30
	// myArray[3] = 40 // it doesn't compile: invalid argument: index 3 out of bounds

	for i, v := range myArray {
		fmt.Printf("i: %d, v: %d\n", i, v)
	}
	/*
	  i: 0, v: 10
	  i: 1, v: 20
	  i: 2, v: 30
	*/
}

func sliceExample() {
	fmt.Println("> sliceExample")
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	printSlice("mySlice", mySlice)

	sliceUntil := mySlice[:4]
	printSlice("sliceUnit", sliceUntil)

	sliceStartingFrom := mySlice[2:]
	printSlice("sliceStartingFrom", sliceStartingFrom)

	sliceFromUntil := mySlice[3:6]
	printSlice("sliceFromUntil", sliceFromUntil)

	// const invalidSliceAsConst = []int{1, 2, 3} // compile error

	var anotherSlice = []int{1, 2, 3}
	printSlice("anotherSlice", anotherSlice)
	anotherSlice = append(anotherSlice, 4)
	printSlice("anotherSlice", anotherSlice)

	predefinedSlice := make([]int, 0, 10)
	printSlice("predefinedSlice", predefinedSlice)

	/*
		mySlice -> length: 10, capacity: 10, value: [1 2 3 4 5 6 7 8 9 10]
		sliceUnit -> length: 4, capacity: 10, value: [1 2 3 4]
		sliceStartingFrom -> length: 8, capacity: 8, value: [3 4 5 6 7 8 9 10]
		sliceFromUntil -> length: 3, capacity: 7, value: [4 5 6]
		anotherSlice -> length: 3, capacity: 3, value: [1 2 3]
		anotherSlice -> length: 4, capacity: 6, value: [1 2 3 4]
		predefinedSlice -> length: 0, capacity: 10, value: []
	*/
}

func printSlice(name string, slice []int) {
	fmt.Printf("%v -> length: %d, capacity: %d, value: %v\n", name, len(slice), cap(slice), slice)
}
