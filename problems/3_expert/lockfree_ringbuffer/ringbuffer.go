package lockfreeringbuffer

import "sync/atomic"

// RingBuffer is a single-producer single-consumer ring buffer (simplified).
type RingBuffer[T any] struct {
	data []T
	size uint64
	head atomic.Uint64 // next index to dequeue
	tail atomic.Uint64 // next index to enqueue
}

func New[T any](capacity uint64) *RingBuffer[T] {
	if capacity == 0 {
		capacity = 1
	}
	return &RingBuffer[T]{data: make([]T, capacity), size: capacity}
}

// Enqueue adds v if buffer not full, returns success.
func (r *RingBuffer[T]) Enqueue(v T) bool {
	for {
		h := r.head.Load()
		t := r.tail.Load()
		if t-h >= r.size {
			return false // full
		}
		if r.tail.CompareAndSwap(t, t+1) {
			idx := t % r.size
			r.data[idx] = v
			return true
		}
	}
}

// Dequeue removes and returns a value if not empty.
func (r *RingBuffer[T]) Dequeue() (v T, ok bool) {
	for {
		h := r.head.Load()
		t := r.tail.Load()
		if h == t {
			var zero T
			return zero, false // empty
		}
		if r.head.CompareAndSwap(h, h+1) {
			idx := h % r.size
			return r.data[idx], true
		}
	}
}
