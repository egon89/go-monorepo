package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	result := Add(2, 2)
	expected := 4

	assertEquals(t, result, expected)
}

// use Example to show how to use the code, which are checked as part of our tests
func ExampleAdd() {
	result := Add(1, 5)
	fmt.Println(result)
	// Output: 6
}

func assertEquals(t testing.TB, result, expected int) {
	t.Helper()

	if result != expected {
		t.Errorf("result: %q. expected: %q", result, expected)
	}
}
