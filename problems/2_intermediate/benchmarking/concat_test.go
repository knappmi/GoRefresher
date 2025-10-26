package benchmarking

import "testing"

var sample = []string{"a","bb","ccc","dddd","ee"}

func BenchmarkConcatPlus(b *testing.B) { for i:=0;i<b.N;i++ { _ = ConcatPlus(sample) } }
func BenchmarkStringsBuilder(b *testing.B) { for i:=0;i<b.N;i++ { _ = ConcatBuilder(sample) } }
