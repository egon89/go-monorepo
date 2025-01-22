package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17}

	target := 7

	result := BinarySearch(arr, target)

	if result != -1 {
		fmt.Printf("Target %d found at index %d\n", target, result)
	} else {
		fmt.Printf("Target %d not found in the array\n", target)
	}
}

func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		}

		if target < arr[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
