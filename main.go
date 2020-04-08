package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	svc := NewCustomerService()

	createCustomerHandler := httptransport.NewServer(
		makeUppercaseEndpoint(svc),
		decodeCreateCustomerRequest,
		encodeResponse,
	)

	http.Handle("/customer", createCustomerHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))
}
