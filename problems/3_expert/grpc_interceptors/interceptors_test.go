package grpcint

import "testing"

func TestChain(t *testing.T) {
	base := func(req any) (any, error) { return req.(int)+1, nil }
	intA := func(req any, next Handler) (any, error) { return next(req.(int)+1) } // +1
	intB := func(req any, next Handler) (any, error) { return next(req.(int)*2) } // *2
	// Chain applies interceptors in provided order wrapping the handler: final order: intA(intB(base))
	h := Chain(base, intA, intB)
	v, _ := h(3) // intB first: 3*2=6, intA: 6+1=7, base: 7+1=8
	if v.(int)!=8 { t.Fatalf("want 8 got %d", v.(int)) }
}
