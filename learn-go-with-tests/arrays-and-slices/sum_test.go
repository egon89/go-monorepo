package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3} // slice

		result := Sum(numbers)
		expected := 6

		if result != expected {
			t.Errorf("result %d, expected %d", result, expected)
		}
	})

}
