package main

import "testing"

func TestPerimeter(t *testing.T) {
	input := Rectangle{10.0, 10.0}
	result := Perimeter(input)
	expected := 40.0

	if result != expected {
		t.Errorf("result %.2f expected %.2f", result, expected)
	}
}

func TestArea(t *testing.T) {
	input := Rectangle{12.0, 6.0}
	result := Area(input)
	expected := 72.0

	if result != expected {
		t.Errorf("result %.2f expected %.2f", result, expected)
	}
}
