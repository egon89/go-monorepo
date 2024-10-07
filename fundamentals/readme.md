# Fundamentals

## Constants and variables
- We can use the short way to declare variables
    ```go
    conferenceName := "Go Conference"
    ```
- We cannot use it
    - for constants
    - for explicity data type like `variableName uint := 2`

## Data Types

### Int
| Go | Java |
| -- | -- |
| int8  | byte  |
| int16 | short |
| int32 | integer |
| int64 | long |

Use `uint8`, `uint16`, `uint32` and `uint64` for positive numbers

## Pointers

### * (Asterisk) -> dereference
- It is used to declare a pointer variable. For example, `*int` represents a pointer to an integer, `*string` represents a pointer to a string, and so on.
- It is also used to dereference a pointer, which means accessing the value stored at the memory address pointed by the pointer variable.

### & (Ampersand):
- It is used to get the memory address of a variable. For example, `&num` gives the memory address of the variable `num`.
