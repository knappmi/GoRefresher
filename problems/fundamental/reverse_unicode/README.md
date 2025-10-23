# Reverse Unicode

Implement a function that reverses a UTF-8 string correctly.

Goals:
- Reverse by Unicode code points.
- Stretch: Reverse by grapheme clusters (use golang.org/x/text/segmenter or rune iteration + combining mark detection).
- Provide a streaming version that works without allocating full slice of runes (two-pass or bytes buffer).

Functions:
- Reverse(s string) string // code point reverse
- ReverseGraphemes(s string) string // stretch (can return s with TODO initially)

Edge Cases:
- Empty string
- ASCII
- Multi-byte runes (e.g., "你好")
- Combining marks (e.g., "Go🇺🇸", "â" (a + combining ^))

Testing:
- Table tests for each category.
- Fuzz test: round trip property len([]rune(s)) == len([]rune(Reverse(Reverse(s))))
