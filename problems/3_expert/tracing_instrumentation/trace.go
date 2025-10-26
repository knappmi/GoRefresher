package tracing

import "time"

// Span represents a tracing span (simplified).
type Span struct { Name string; Start time.Time; End time.Time }

var spans []Span

// StartSpan starts a span and returns a function to end it.
func StartSpan(name string) func() {
	sp := Span{Name:name, Start: time.Now()}
	idx := len(spans)
	spans = append(spans, sp)
	return func(){ spans[idx].End = time.Now() }
}

// Spans returns recorded spans.
func Spans() []Span { return spans }
