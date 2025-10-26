package reverseunicode

// Reverse returns the input reversed by code points.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// ReverseGraphemes returns input reversed by grapheme clusters (fallback to rune reversal).
func ReverseGraphemes(s string) string {
	// Simplified: treat each rune as a grapheme.
	return Reverse(s)
}
