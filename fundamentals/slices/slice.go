package slices

import "fmt"

func SliceExample() {
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
