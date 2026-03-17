// leaf service - no downstream calls from here
package servicec

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kiing-dom/phylax/pkg/tracing"
)

func main() {
	exporter := tracing.NewHTTPExporter("http://localhost:8080")
	tracer := tracing.NewTracer("service-c", exporter)

	http.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
		ctx := tracing.Extract(r)

		var span *tracing.Span

		// if not part of a trace (root span)
		if ctx.TraceID == "" {
			span = tracer.StartSpan("service-c /work")
		} else {
			span = tracer.StartSpanFromContext(ctx, "service-c /work")
		}

		defer span.End()

		// simulate an actual task (/work)
		fmt.Println(w, "service C response")
	})

	log.Println("service c running on port :8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
}
