# Lock-Free Ring Buffer

Implement a multi-producer, multi-consumer (MPMC) ring buffer using atomics.

Goals:
- Generic type parameter T any.
- Non-blocking Enqueue/Dequeue that return bool success.
- Use power-of-two capacity for masking.
- Compare against channel performance (benchmark).

Stretch:
- Support blocking variant with runtime.Gosched spin then backoff.
