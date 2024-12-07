package main

import (
	"fmt"
	"math"
)

type MyFloat float64

// numeric type MyFloat with an Abs method
func (f MyFloat) Abs() float64 {
	if f < 0 {
		// negating -f makes -1.41 positive
		// --1.41 = 1.41
		return float64(-f)
	}

	return float64(f)
}

func main() {
	fmt.Println("methods continued")
	f := MyFloat(-math.Sqrt2) // -1.41
	fmt.Println(f.Abs())
}
