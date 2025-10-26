package goroutinesbasics

import "sync/atomic"

// IncrementUnsafe increments a shared counter without synchronization (intentionally racy under concurrency).
func IncrementUnsafe(counter *int64, times int) {
	for i := 0; i < times; i++ {
		*counter = *counter + 1
	}
}

// IncrementAtomic increments using atomic operations.
func IncrementAtomic(counter *int64, times int) {
	for i := 0; i < times; i++ {
		atomic.AddInt64(counter, 1)
	}
}
