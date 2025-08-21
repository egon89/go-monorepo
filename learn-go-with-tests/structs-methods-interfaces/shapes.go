package main

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

/*
When your method is called on a variable of that type,
you get your reference to its data via the receiverName variable.
In many other programming languages this is done implicitly and you access the receiver via this.
*/
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
