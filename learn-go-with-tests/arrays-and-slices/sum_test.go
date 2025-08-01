package main

import (
	"slices"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	result := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	if !slices.Equal(result, expected) {
		t.Errorf("result %v, expected %v", result, expected)
	}
}

func TestSumAllTails(t *testing.T) {
	result := SumAllTails([]int{1, 2}, []int{0, 9}, []int{3}, []int{})
	expected := []int{2, 9, 0, 0}

	if !slices.Equal(result, expected) {
		t.Errorf("result %v, expected %v", result, expected)
	}
}
