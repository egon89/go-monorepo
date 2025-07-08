package iteration

import "testing"

func TestRepeat(t *testing.T) {
	result := Repeat("a")
	expected := "aaaaa"

	if result != expected {
		t.Errorf("result: %q. expected: %q", result, expected)
	}
}

func TestRepeatStringBuilder(t *testing.T) {
	result := Repeat("b")
	expected := "bbbbb"

	if result != expected {
		t.Errorf("result: %q. expected: %q", result, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	// we can use b.Loop in go 1.24+

	b.ResetTimer() // Reset timer to exclude setup time
	for i := 0; i < b.N; i++ {
		// Code to be benchmarked
		Repeat("a")
	}
}

func BenchmarkRepeatStringBuilder(b *testing.B) {
	// we can use b.Loop in go 1.24+

	b.ResetTimer() // Reset timer to exclude setup time
	for i := 0; i < b.N; i++ {
		// Code to be benchmarked
		RepeatStringBuilder("b")
	}
}
