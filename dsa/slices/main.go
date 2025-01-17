package main

import "fmt"

/*
While arrays have a fixed size, slices are dynamic and can grow or shrink as needed.
They are built on top of arrays but provide additional functionality.
A slice is a reference to a portion of an array. Unlike arrays, slices do not have a fixed size
*/
func main() {
	fmt.Println("> slices")

	createASliceFromArray()
	createASliceDirectly()
	createALiteralSlice()
	copyElements()
	sliceASlice()
}

func createASliceFromArray() {
	fmt.Println("> createASliceFromArray")

	arr := [5]int{10, 20, 30, 40, 50}

	slice := arr[1:4] // from index 1 to 3 (4 is excluded)

	fmt.Println("Slice:", slice)         // Slice: [20 30 40]
	fmt.Println("Length:", len(slice))   // Length: 3
	fmt.Println("Capacity:", cap(slice)) // Capacity: 4
}

func createASliceDirectly() {
	fmt.Println("> createASliceDirectly")

	slice := make([]int, 3, 5) // length: 3, capacity: 5

	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	// slice[3] = 4 // error

	fmt.Println("Slice:", slice)         // Slice: [1 2 3]
	fmt.Println("Length:", len(slice))   // Length: 3
	fmt.Println("Capacity:", cap(slice)) // Capacity: 5

	slice = append(slice, 4)

	fmt.Println("Slice:", slice)         // Slice: [1 2 3 4]
	fmt.Println("Length:", len(slice))   // Length: 4
	fmt.Println("Capacity:", cap(slice)) // Capacity: 5

	slice = append(slice, 5)
	slice = append(slice, 6)

	fmt.Println("Slice:", slice)         // Slice: [1 2 3 4 5 6]
	fmt.Println("Length:", len(slice))   // Length: 6
	fmt.Println("Capacity:", cap(slice)) // Capacity: 10 (previous capacity * 2)
}

func createALiteralSlice() {
	fmt.Println("> createALiteralSlice")

	slice := []int{5, 10, 15, 20}

	fmt.Println("Slice:", slice)         // Slice: [5 10 15 20]
	fmt.Println("Length:", len(slice))   // Length: 4
	fmt.Println("Capacity:", cap(slice)) // Capacity: 4
}

func copyElements() {
	fmt.Println("> copyElements")

	source := []int{10, 20, 30}
	dest := make([]int, 2)

	copy(dest, source)

	fmt.Println("source:", source)    // source: [10 20 30]
	fmt.Println("destination:", dest) // destination: [10 20]
}

func sliceASlice() {
	fmt.Println("> sliceASlice")

	slice := []int{1, 2, 3, 4, 5, 6}
	partialSlice := slice[1:4]

	fmt.Println("partial slice:", partialSlice) // [2 3 4]

	partialSliceAfter := slice[2:]
	fmt.Println("partial slice after:", partialSliceAfter) // [3 4 5 6]

	partialSliceBefore := slice[:5]
	fmt.Println("partial slice after:", partialSliceBefore) // [1 2 3 4 5]

}
