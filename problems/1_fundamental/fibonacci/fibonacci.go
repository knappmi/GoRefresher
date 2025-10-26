package fibonacci

// FibRecursive returns nth Fibonacci using naive recursion.
func FibRecursive(n int) int {
	if n < 2 {
		return n
	}
	return FibRecursive(n-1) + FibRecursive(n-2)
}

// FibIterative returns nth Fibonacci using loop.
func FibIterative(n int) int {
	if n < 2 {
		return n
	}
	prev, curr := 0, 1
	for i := 2; i <= n; i++ {
		prev, curr = curr, prev+curr
	}
	return curr
}

// FibMemoized uses provided cache map for memoization.
func FibMemoized(n int, cache map[int]int) int {
	if n < 2 {
		return n
	}
	if v, ok := cache[n]; ok {
		return v
	}
	v := FibMemoized(n-1, cache) + FibMemoized(n-2, cache)
	cache[n] = v
	return v
}
