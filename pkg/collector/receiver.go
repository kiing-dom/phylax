package collector

import (
	"encoding/json"
	"net/http"

	"github.com/kiing-dom/phylax/pkg/tracing"
)

type Receiver struct {
	storage *Storage
}

func NewReceiver(storage *Storage) *Receiver {
	return &Receiver{
		storage: storage,
	}
}

func (r *Receiver) HandleSpans(w http.ResponseWriter, req *http.Request) {
	var span tracing.Span

	err := json.NewDecoder(req.Body).Decode(&span)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	r.storage.AddSpan(span)

	w.WriteHeader(http.StatusAccepted)
}
