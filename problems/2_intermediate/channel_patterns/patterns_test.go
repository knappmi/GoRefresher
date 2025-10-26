package channelpatterns

import "testing"

func TestFanInBasic(t *testing.T) {
	c1 := make(chan int, 2); c2 := make(chan int, 2)
	c1<-1; c1<-2; close(c1)
	c2<-3; c2<-4; close(c2)
	out := FanIn(c1,c2)
	res := make(map[int]bool)
	for v := range out { res[v] = true }
	for _, v := range []int{1,2,3,4} { if !res[v] { t.Fatalf("missing %d", v) } }
}
