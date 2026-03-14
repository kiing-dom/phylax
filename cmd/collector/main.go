package main

import (
	"log"
	"net/http"

	"github.com/kiing-dom/phylax/pkg/collector"
)

func main() {
	storage := collector.NewStorage()
	receiver := collector.NewReceiver(storage)

	http.HandleFunc("/spans", receiver.HandleSpans)
	log.Println("Collector running on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
