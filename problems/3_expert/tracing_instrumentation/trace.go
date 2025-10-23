package tracing

// StartSpan starts a span and returns end func.
func StartSpan(name string) func() { // TODO: integrate OpenTelemetry
	return func(){} }
