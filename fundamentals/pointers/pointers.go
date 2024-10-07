package pointers

import "fmt"

func PointerExample() {
	fmt.Println("> pointerExample")
	a := 10
	fmt.Printf("a: %v\n", a)
	// * (asterisk) -> dereference
	var aPointer *int = &a
	fmt.Printf("aPointer: %v\n", aPointer)

	*aPointer = 20
	fmt.Printf("a: %v\n", a)
	b := &a
	fmt.Printf("b as pointer: %v, b as pointer value using *: %v\n", b, *b)

	*b = 90
	fmt.Printf("a: %v, *aPointer: %v, *b: %v\n", a, *aPointer, *b)
	fmt.Printf("&a: %v, aPointer: %v, b: %v\n", &a, aPointer, b)

	/*
		a: 10
		aPointer: 0xc00009a1d0
		a: 20
		b as pointer: 0xc00009a1d0, b as pointer value using *: 20
		a: 90, *aPointer: 90, *b: 90
		&a: 0xc00009a1d0, aPointer: 0xc00009a1d0, b: 0xc00009a1d0
	*/
}
