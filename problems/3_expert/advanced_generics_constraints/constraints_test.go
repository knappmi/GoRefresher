package advconstraints

import "testing"

func TestAdd(t *testing.T) { if Add(2,3)!=5 || Add(float64(1.5), float64(2.5))!=4.0 { t.Fatal("bad add") } }
