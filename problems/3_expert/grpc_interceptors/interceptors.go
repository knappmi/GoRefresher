package grpcint

// TODO: implement interceptor chain with gRPC

type Handler func(req any) (any, error)

type Interceptor func(req any, next Handler) (any, error)

// Chain composes interceptors so that ints[0] is outermost and ints[last] innermost before handler.
func Chain(h Handler, ints ...Interceptor) Handler {
	for i := 0; i < len(ints); i++ {
		intc := ints[i]
		next := h
		h = func(req any) (any, error) { return intc(req, next) }
	}
	return h
}
