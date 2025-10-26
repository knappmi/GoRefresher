package profiling

import "testing"

func BenchmarkNaiveSum(b *testing.B) { for i:=0;i<b.N;i++ { _ = NaiveSum(10000) } }
func BenchmarkFastSum(b *testing.B) { for i:=0;i<b.N;i++ { _ = FastSum(10000) } }
