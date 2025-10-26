package racer

import "testing"
import "sync"

func TestRaceSafe(t *testing.T) {
	shared = 0
	wg := sync.WaitGroup{}
	for i:=0;i<1000;i++ { wg.Add(1); go func(){ Inc(); wg.Done() }() }
	wg.Wait()
	if Value()!=1000 { t.Fatalf("want 1000 got %d", Value()) }
}

func TestRaceUnsafe(t *testing.T) { t.Skip("Run with -race to observe data race and wrong final value") }
