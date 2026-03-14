`pkg/tracing/span.go`

core data model: **spans**
**span**: a timed record of an operation

**trace**: a collection of spans
e.g.

```plaintext
Trace 123

Span A (service A)
   └─ Span B (service B)
         └─ Span C (service C)

```

### modelling a span
```go
TraceID string
SpanID string
ParentSpanID string

Service string
Operation string

StartTime time.Time
Duration time.Duration
```

