package lockfreeringbuffer

// TODO: implement lock-free ring buffer

type RingBuffer[T any] struct{}

func New[T any](capacity uint64) *RingBuffer[T] { // TODO: allocate internal structure
	return &RingBuffer[T]{} }
func (r *RingBuffer[T]) Enqueue(v T) bool { // TODO implement
	return false }
func (r *RingBuffer[T]) Dequeue() (v T, ok bool) { // TODO implement
	var zero T; return zero, false }
