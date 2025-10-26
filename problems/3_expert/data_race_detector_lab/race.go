package racer

import "sync/atomic"

var shared int64

// Inc performs a safe atomic increment.
func Inc() { atomic.AddInt64(&shared, 1) }
// UnsafeInc performs a non-atomic increment (intentional race under concurrency).
func UnsafeInc() { shared = shared + 1 }
func Value() int64 { return shared }
