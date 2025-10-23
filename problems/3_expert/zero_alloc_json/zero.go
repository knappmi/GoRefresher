package zeroallocjson

import "sync"

var bufPool = sync.Pool{ New: func() any { b := make([]byte, 0, 1024); return &b } }

// TODO: implement zero-allocation JSON marshal
func Marshal(v any) []byte { return []byte("{}") }
