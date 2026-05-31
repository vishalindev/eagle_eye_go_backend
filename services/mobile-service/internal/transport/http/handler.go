package http

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"

	"example.com/eagle-eye/services/mobile-service/internal/endpoints"
)

// NewHandler exposes health and endpoint-registry routes over HTTP.
func NewHandler(endpointSet endpoints.Set) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("GET /healthz", httptransport.NewServer(endpointSet.Health, decodeEmptyRequest, encodeJSONResponse))
	mux.Handle("GET /endpoints", httptransport.NewServer(endpointSet.Endpoints, decodeEmptyRequest, encodeJSONResponse))
	return mux
}

func decodeEmptyRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return struct{}{}, nil
}

func encodeJSONResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
