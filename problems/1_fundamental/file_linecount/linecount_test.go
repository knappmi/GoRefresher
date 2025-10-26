package linecount

import (
	"os"
	"testing"
)

func TestCountLines(t *testing.T) {
	fname := t.TempDir() + "/test.txt"
	if err := os.WriteFile(fname, []byte("a\n b\n c"), 0644); err != nil {
		t.Fatal(err)
	}
	c, err := CountLines(fname)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if c != 3 {
		t.Fatalf("want 3 got %d", c)
	}
}
