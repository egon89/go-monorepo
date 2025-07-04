package main

import "testing"

func TestHello(t *testing.T) {
	result := Hello("John")
	expected := "Hello John"

	if result != expected {
		t.Errorf("result: %q. expected: %q", result, expected)
	}
}
