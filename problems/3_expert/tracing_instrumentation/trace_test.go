package tracing

import "testing"

func TestSpan(t *testing.T) {
	spans = nil
	end := StartSpan("x")
	end()
	if len(Spans())!=1 || Spans()[0].End.IsZero() { t.Fatal("span not recorded") }
}
