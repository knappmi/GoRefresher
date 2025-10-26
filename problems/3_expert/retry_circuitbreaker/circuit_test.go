package circuit

import "testing"

func TestBreaker(t *testing.T) {
	b := New()
	if !b.Allow() { t.Fatal("should allow") }
	b.Fail(); b.Fail(); b.Fail()
	if b.Allow() { t.Fatal("should be open") }
}

func TestBackoff(t *testing.T) { if Backoff(3, 10)!=80 { t.Fatal("bad backoff") } }
