package conditionals

import (
	"fmt"
	"math"
)

func IfWithShortStatementToExecuteBefore() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	/*
		9 < 10, returning v 9
		27 >= 20, returning limit 20
		9 20
	*/
}

func pow(x, y, limit float64) float64 {
	if v := math.Pow(x, y); v < limit {
		fmt.Printf("%v < %v, returning v %v\n", v, limit, v)
		return v
	} else {
		fmt.Printf("%v >= %v, returning limit %v\n", v, limit, limit)
	}
	// v is not available here!

	return limit
}
