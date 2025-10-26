package genericscollections

import "testing"

func TestSet(t *testing.T) {
	s := NewSet[int](); s.Add(1); s.Add(2); if !s.Has(1) || !s.Has(2) { t.Fatal("missing") }
	s.Remove(1); if s.Has(1) { t.Fatal("remove failed") }
}
