package wordcount

import (
	"bufio"
	"io"
	"unicode"
	"strings"
)

// Count returns word frequencies from reader.
func Count(r io.Reader) (map[string]int, error) {
	scanner := bufio.NewScanner(r)
	freq := make(map[string]int)
	for scanner.Scan() {
		line := scanner.Text()
		var b strings.Builder
		flush := func(){ if b.Len()>0 { freq[b.String()]++; b.Reset() } }
		for _, rn := range line {
			if unicode.IsLetter(rn) || rn=='\'' { // allow apostrophes inside words
				b.WriteRune(unicode.ToLower(rn))
			} else { flush() }
		}
		flush()
	}
	return freq, scanner.Err()
}
