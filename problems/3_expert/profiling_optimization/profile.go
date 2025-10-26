package profiling

// NaiveSum squares then sums numbers.
func NaiveSum(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i * i
	}
	return sum
}

// FastSum uses formula sum_{i=0}^{n-1} i^2 = (n-1)*n*(2n-1)/6
func FastSum(n int) int {
	if n <= 0 {
		return 0
	}
	m := n - 1
	return (m * n * (2*n - 1)) / 6
}
