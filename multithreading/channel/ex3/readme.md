# Multithreading Example: Channel Synchronization

## Overview

This example demonstrates how to use channels and goroutines in Go to synchronize the execution of tasks. The program processes pairs of values (even and odd) using semaphores to ensure that the "even" task completes before the corresponding "odd" task starts.

## Code Explanation

### Main Function

The `main` function orchestrates the execution of the program:

1. **Initialization**:
   - A `sync.WaitGroup` is used to wait for all goroutines to complete.
   - A slice of channels (`semaphores`) is created to act as semaphores for each pair of tasks.

   ```go
   var wg sync.WaitGroup
   pairs := 5 // 0-1, 2-3, 4-5, 6-7, 8-9
   semaphores := make([]chan struct{}, pairs)

   for i := 0; i < pairs; i++ {
       semaphores[i] = make(chan struct{})
   }
   ```

2. **Processing Pairs**:
   - For each pair of values (even and odd):
     - A goroutine is launched to process the "even" value.
     - Another goroutine is launched to process the "odd" value, which waits for the "even" value to complete.

   ```go
   for i := 0; i < pairs; i++ {
       even := i * 2
       odd := even + 1

       wg.Add(1)
       go func(value int, channel chan struct{}) {
           defer wg.Done()
           fmt.Printf("# working on even value %d\n", value)
           sleep(value)
           fmt.Printf("## even value completed: %d\n", value)
           channel <- struct{}{} // release odd
       }(even, semaphores[i])

       wg.Add(1)
       go func(value int, channel chan struct{}) {
           defer wg.Done()
           fmt.Printf("# waiting for even value %d\n", even)
           <-channel // wait for even to complete
           fmt.Printf("## working on odd value %d [pair %d-%d] \n", value, even, value)
       }(odd, semaphores[i])
   }
   ```

3. **Waiting for Completion**:
   - The `WaitGroup` ensures that the program waits for all goroutines to finish before exiting.

   ```go
   wg.Wait()
   fmt.Println("...done")
   ```

### Sleep Function

The `sleep` function simulates work by pausing execution for a duration proportional to the input value.

```go
func sleep(value int) {
    duration := time.Duration(value) * time.Millisecond
    fmt.Printf("...value %d will sleep %v ms\n", value, duration)
    time.Sleep(duration)
}
```

## Output Example

When you run the program, you will see output similar to the following:

```
# working on even value 0
...value 0 will sleep 0ms
## even value completed: 0
# waiting for even value 0
## working on odd value 1 [pair 0-1]
# working on even value 2
...value 2 will sleep 2ms
## even value completed: 2
# waiting for even value 2
## working on odd value 3 [pair 2-3]
...
...done
```

## Key Concepts

- **Channels**: Used as semaphores to synchronize the execution of tasks.
- **Goroutines**: Lightweight threads used to execute tasks concurrently.
- **WaitGroup**: Ensures that the program waits for all goroutines to complete before exiting.

## How to Run

1. Navigate to the directory containing the `main.go` file.
2. Run the program using the following command:

   ```bash
   go run main.go
   ```
