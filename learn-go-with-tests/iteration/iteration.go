package iteration

import "strings"

const repeatCount = 5

func Repeat(s string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += s
	}

	return repeated
}

func RepeatStringBuilder(s string) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(s)
	}

	return repeated.String()
}
