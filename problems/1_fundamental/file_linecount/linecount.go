package linecount

import (
	"bufio"
	"io"
	"os"
)

// CountLines counts lines in a file path.
func CountLines(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	count := 0
	empty := true
	for {
		line, err := reader.ReadBytes('\n')
		if len(line) > 0 {
			empty = false
			count++
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return count, err
		}
	}
	if empty {
		return 0, nil
	}
	return count, nil
}
