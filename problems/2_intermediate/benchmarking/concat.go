package benchmarking

import "strings"

// ConcatPlus concatenates using '+' in a loop.
func ConcatPlus(parts []string) string {
	out := ""
	for _, p := range parts {
		out += p
	}
	return out
}

// ConcatBuilder uses strings.Builder.
func ConcatBuilder(parts []string) string {
	var b strings.Builder
	for _, p := range parts {
		b.WriteString(p)
	}
	return b.String()
}
