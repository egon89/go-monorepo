# Arrays and Slices
In Go, arrays are fixed-size sequences of elements of the same type, while slices are dynamically-sized, flexible views into the elements of an array. Slices are more commonly used due to their versatility.

```go
// Array
var a [5]int
a[0] = 1

// Slice
s := []int{1, 2, 3}
s = append(s, 4) // Slices can grow dynamically
```

## Variadic Functions

In Go, [variadic functions](https://gobyexample.com/variadic-functions) allow you to pass a variable number of arguments to a function. This is useful when you want to handle an unknown number of inputs without defining multiple function signatures.

```go
func sum(nums ...int) {
    // nums is a slice of int
}
```
