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

## Go Slice: Length vs Capacity

### 📌 Quick Recap

- `len(slice)` → Number of elements **present** in the slice.
- `cap(slice)` → Number of elements the slice can **hold** in memory before reallocation.


### 🧪 Example code

```go
func main() {
	// Create a slice with length 0 and capacity 5
	s := make([]int, 0, 5)
	fmt.Println("Initial slice:")
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(s), cap(s), s)

	// Append one item
	s = append(s, 10)
	fmt.Println("\nAfter appending 1 element:")
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(s), cap(s), s)

	// Append more items up to capacity
	s = append(s, 20, 30, 40, 50)
	fmt.Println("\nAfter appending up to capacity:")
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(s), cap(s), s)

	// Append beyond the initial capacity
	s = append(s, 60)
	fmt.Println("\nAfter appending beyond capacity:")
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(s), cap(s), s)
}
```

### 🧪 Output

```
Initial slice:
len=0, cap=5, slice=[]

After appending 1 element:
len=1, cap=5, slice=[10]

After appending up to capacity:
len=5, cap=5, slice=[10 20 30 40 50]

After appending beyond capacity:
len=6, cap=10, slice=[10 20 30 40 50 60]
```

### 🧠 What's happening?
- Initially, the slice has 0 elements but room for 5, so len=0, cap=5.
- Appending up to 5 elements just fills the preallocated space.
- Once you go beyond 5 elements, Go internally creates a **new array** (usually doubling the capacity) and copies existing elements to it. So now cap=10.

### ✅ Summary:
- **Length** is how many elements you're actually using.
- **Capacity** is how much memory Go has already reserved for you.
- **make([]T, length, capacity)** gives you control over both, useful for performance (e.g., avoiding many reallocations).

