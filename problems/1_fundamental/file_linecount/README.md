# File Line Count

Implement efficient line counting for large files.

Function:
- CountLines(path string) (int, error)

Rules:
- Use bufio.Reader and count '\n'.
- Handle last line without newline.

Stretch:
- Concurrent chunked counting (memory map or splits) comparing speed.
- Support stdin when path == "-".
