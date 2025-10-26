package jsonstreaming

import (
	"encoding/json"
	"io"
)

// DecodeStream processes a JSON array from reader returning element count.
func DecodeStream(r io.Reader) (int, error) {
	dec := json.NewDecoder(r)
	// Expect start of array
	tok, err := dec.Token(); if err!=nil { return 0, err }
	if delim, ok := tok.(json.Delim); !ok || delim != '[' { return 0, &json.SyntaxError{Offset:0} }
	count := 0
	for dec.More() {
		var v interface{}
		if err := dec.Decode(&v); err!=nil { return count, err }
		count++
	}
	// consume closing ]
	_, err = dec.Token(); if err!=nil { return count, err }
	return count, nil
}
