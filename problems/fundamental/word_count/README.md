# Word Count

Implement a word frequency counter.

Function:
- Count(r io.Reader) (map[string]int, error)

Rules:
- Normalize to lower case.
- Split on non-letter (unicode) boundaries.
- Ignore empty tokens.

Add CLI main.go to read from file arg or stdin.

Stretch:
- Top N words output.
- Concurrency: process chunks.
- Stream large file with limited memory.
