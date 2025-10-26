package customscheduler

import "testing"

func TestScheduler(t *testing.T) {
	s := New(); total := 0
	for i:=0;i<100;i++ { s.Submit(func(){ total++ }) }
	s.Close(); s.Run(4)
	if total != 100 { t.Fatalf("want 100 got %d", total) }
}
