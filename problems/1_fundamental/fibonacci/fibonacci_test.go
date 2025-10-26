package fibonacci

import "testing"

func TestFibBasic(t *testing.T) {
	cases := []struct{ n, want int }{{0,0},{1,1},{2,1},{3,2},{10,55},{20,6765}}
	for _, c := range cases {
		if got := FibIterative(c.n); got != c.want { t.Fatalf("Iterative %d => %d want %d", c.n, got, c.want) }
		if got := FibRecursive(c.n); got != c.want { t.Fatalf("Recursive %d => %d want %d", c.n, got, c.want) }
		if got := FibMemoized(c.n, map[int]int{}); got != c.want { t.Fatalf("Memoized %d => %d want %d", c.n, got, c.want) }
	}
}

func BenchmarkFibIterative(b *testing.B) { for i:=0;i<b.N;i++ { _ = FibIterative(30) } }
func BenchmarkFibRecursive(b *testing.B) { for i:=0;i<b.N;i++ { _ = FibRecursive(30) } }
func BenchmarkFibMemoized(b *testing.B) { for i:=0;i<b.N;i++ { _ = FibMemoized(30, map[int]int{}) } }
