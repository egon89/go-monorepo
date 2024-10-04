package main

import (
	"fmt"
	"fundamentals/arrays"
	"fundamentals/functions"
	"fundamentals/maps"
	"fundamentals/slices"
)

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
	arrays.ArrayExample()
	slices.SliceExample()
	maps.MapExample()
	functions.FunctionExample()
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
