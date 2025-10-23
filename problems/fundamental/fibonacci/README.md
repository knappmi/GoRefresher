# Fibonacci Implementations

Implement Fibonacci three ways and compare performance.

Functions:
- FibRecursive(n int) int // naive recursion
- FibIterative(n int) int // loop
- FibMemoized(n int, cache map[int]int) int // top-down memoization

Tasks:
1. Implement each.
2. Add benchmarks (go test -bench .) and compare allocs.
3. Discuss when recursion vs iteration is acceptable.

Stretch:
- Use big.Int for large n.
- Add iterative with matrix exponentiation or fast doubling.
