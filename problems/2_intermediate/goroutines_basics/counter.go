package goroutinesbasics

// IncrementUnsafe increments a shared counter without synchronization.
func IncrementUnsafe(counter *int64, times int) {
	// TODO: implement intentionally unsafe increment
}

// IncrementAtomic increments using atomic operations.
func IncrementAtomic(counter *int64, times int) {
	// TODO: implement atomic increment
}
