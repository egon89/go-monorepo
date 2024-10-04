package arrays

import "fmt"

func ArrayExample() {
	fmt.Println("> arrayExample")
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
