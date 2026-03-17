package tracing

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type HTTPExporter struct {
	CollectorURL string
}

func NewHTTPExporter(url string) *HTTPExporter {
	return &HTTPExporter{
		CollectorURL: url,
	}
}

func (e *HTTPExporter) Export(span *Span) error {
	data, err := json.Marshal(span)
	if err != nil {
		return err
	}

	_, err = http.Post(
		e.CollectorURL+"/spans",
		"application/json",
		bytes.NewBuffer(data),
	)

	return err
}
