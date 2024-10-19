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
	changeValue()
}

func printSlice(name string, slice []int) {
	fmt.Printf("%v -> length: %d, capacity: %d, value: %v\n", name, len(slice), cap(slice), slice)
}

// slice doesn't store data, it is just a reference to an array
func changeValue() {
	fmt.Println("> changeValue")
	var colors = [4]string{"blue", "red", "yellow", "white"}
	filter1 := colors[1:4]
	fmt.Println("filter1:", filter1) // filter1: [red yellow white]
	filter2 := colors[0:3]
	fmt.Println("filter2:", filter2) // filter1: [blue red yellow]

	filter2[2] = "gray"              // this change is made in the colors array and will be reflected in all slices
	fmt.Println("filter1:", filter1) // filter1: [red gray white]
	fmt.Println("filter2:", filter2) // filter1: [blue red gray]

}
