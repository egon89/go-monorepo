package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("should say hello to people", func(t *testing.T) {
		result := Hello("John")
		expected := "Hello John"

		assertCorrectMessage(t, result, expected)
	})

	t.Run("should say 'Hello world' when an empty string is supplied", func(t *testing.T) {
		result := Hello("")
		expected := "Hello world"

		assertCorrectMessage(t, result, expected)
	})
}

// testing.TB which is an interface that *testing.T and *testing.B both satisfy,
// so you can call helper functions from a test, or a benchmark
func assertCorrectMessage(t testing.TB, result, expected string) {
	// when a test fails, the line number reported will be in our function call rather than inside our test helper
	t.Helper()

	if result != expected {
		t.Errorf("result: %q. expected: %q", result, expected)
	}
}
