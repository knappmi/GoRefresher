package goroutinesbasics

import (
	"runtime"
	"sync"
	"testing"
)

func TestIncrementAtomic(t *testing.T) {
	var c int64
	wg := sync.WaitGroup{}
	workers := runtime.NumCPU()
	wg.Add(workers)
	for i:=0;i<workers;i++ { go func(){ IncrementAtomic(&c, 1000); wg.Done() }() }
	wg.Wait()
	want := int64(workers*1000)
	if c != want { t.Fatalf("want %d got %d", want, c) }
}

func TestIncrementUnsafe(t *testing.T) { t.Skip("Run with -race to see data race and incorrect count") }
