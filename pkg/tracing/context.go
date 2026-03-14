package tracing

type TraceContext struct {
	TraceID      string
	ParentSpanID string
}
