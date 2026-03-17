package tracing

type Exporter interface {
	Export(s *Span) error
}
