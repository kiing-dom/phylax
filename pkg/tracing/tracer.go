package tracing

import "time"

type Tracer struct {
	Service  string
	Exporter Exporter
}

func NewTracer(service string, exporter Exporter) *Tracer {
	return &Tracer{
		Service:  service,
		Exporter: exporter,
	}
}

func (t *Tracer) StartSpan(operation string) *Span {
	return &Span{
		TraceID:      generateID(),
		SpanID:       generateID(),
		ParentSpanID: generateID(),

		Service:   t.Service,
		Operation: operation,

		StartTime: time.Now(),
		exporter:  t.Exporter,
	}
}

func (t *Tracer) StartSpanFromContext(ctx TraceContext, operation string) *Span {
	return &Span{
		TraceID:      ctx.TraceID,
		SpanID:       generateID(),
		ParentSpanID: ctx.ParentSpanID,

		Service:   t.Service,
		Operation: operation,

		StartTime: time.Now(),
		exporter:  t.Exporter,
	}
}
