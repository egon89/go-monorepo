# Map

The **key** type must be **comparable**, meaning we can use `==` and `!=` operators with it ([Comparison operators](https://go.dev/ref/spec#Comparison_operators)). The value type can be any type, including another map or a slice.

> A map value is a pointer to a runtime.hmap structure.

So when you pass a map to a function/method, you are indeed copying it, but just the pointer part, not the underlying data structure that contains the data.

We should never initialize a nil map variable:

```go
var m map[string]string // m is nil
m["route"] = "1"        // panic: assignment to entry in nil map
```

Instead, we should initialize it with `make` or an empty map:

```go
m := make(map[string]string)
// or
m := map[string]string{}
```

## Add elements
Map will not throw an error if the value already exists. Instead, they will go ahead and overwrite the value with the newly provided value.

```go
m := make(map[string]string)
m["route"] = "1"
m["route"] = "2" // no error, value is overwritten
```
