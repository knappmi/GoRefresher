package grpcint

// TODO: implement interceptor chain with gRPC

type Handler func(req any) (any, error)

type Interceptor func(req any, next Handler) (any, error)

func Chain(h Handler, ints ...Interceptor) Handler {
	// TODO: implement chaining logic
	return h
}
