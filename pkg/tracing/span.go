package tracing

import "time"

type Span struct {
	TraceID      string `json:"trace_id"`
	SpanID       string `json:"span_id"`
	ParentSpanID string `json:"parent_span_id,omitempty"`

	Service   string `json:"service"`
	Operation string `json:"operation"`

	StartTime time.Time     `json:"start_time"`
	Duration  time.Duration `json:"duration"`
}

func (s *Span) Context() TraceContext {
	return TraceContext{
		TraceID:      s.TraceID,
		ParentSpanID: s.SpanID,
	}
}

func (s *Span) End() {
	s.Duration = time.Since(s.StartTime)
}
