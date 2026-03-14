package collector

import (
	"sync"

	"github.com/kiing-dom/phylax/pkg/tracing"
)

type Storage struct {
	mu     sync.Mutex
	traces map[string][]tracing.Span // map[TraceId][]Span
}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) AddSpan(span tracing.Span) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.traces[span.TraceID] = append(s.traces[span.TraceID], span)
}

func (s *Storage) GetTrace(traceID string) []tracing.Span {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.traces[traceID]
}
