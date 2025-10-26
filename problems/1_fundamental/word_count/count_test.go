package wordcount

import "strings"
import "testing"

func TestCount(t *testing.T) {
	input := "Hello, hello world! It's a world." // world appears twice
	freq, err := Count(strings.NewReader(input))
	if err!=nil { t.Fatalf("err: %v", err) }
	if freq["hello"]!=2 { t.Fatalf("hello count %d", freq["hello"]) }
	if freq["world"]!=2 { t.Fatalf("world count %d", freq["world"]) }
	if freq["it's"]!=1 { t.Fatalf("apostrophe word missing: %#v", freq) }
}
