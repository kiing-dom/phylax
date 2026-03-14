package tracing

import "net/http"

type TraceContext struct {
	TraceID      string
	ParentSpanID string
}

const (
	TraceIDHeader    = "X-Trace-ID"
	ParentSpanHeader = "X-Parent-Span-ID"
)

func Inject(ctx TraceContext, req *http.Request) {
	req.Header.Set(TraceIDHeader, ctx.TraceID)
	req.Header.Set(ParentSpanHeader, ctx.ParentSpanID)
}
