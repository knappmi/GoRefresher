package jsonstreaming

import "testing"
import "strings"

func TestDecodeStream(t *testing.T) {
	jsonData := `[1,2,3,{"a":4}]`
	count, err := DecodeStream(strings.NewReader(jsonData))
	if err!=nil || count!=4 { t.Fatalf("count %d err %v", count, err) }
}
