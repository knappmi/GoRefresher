package lockfreeringbuffer

import "testing"

func TestBasic(t *testing.T) {
	r := New[int](8)
	for i:=0;i<8;i++ { if !r.Enqueue(i) { t.Fatal("enqueue fail") } }
	if r.Enqueue(99) { t.Fatal("should be full") }
	for i:=0;i<8;i++ { v, ok := r.Dequeue(); if !ok || v!=i { t.Fatalf("bad dequeue %d %v", i, v) } }
	if _, ok := r.Dequeue(); ok { t.Fatal("should be empty") }
}

func BenchmarkEnqueueDequeue(b *testing.B) {
	r := New[int](1024)
	for i:=0;i<b.N;i++ { r.Enqueue(i); r.Dequeue() }
}
