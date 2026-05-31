package http

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

// DecodeRequestFunc extracts an endpoint request from an HTTP request.
type DecodeRequestFunc func(context.Context, *http.Request) (interface{}, error)

// EncodeResponseFunc writes an endpoint response to an HTTP response.
type EncodeResponseFunc func(context.Context, http.ResponseWriter, interface{}) error

// Server adapts a go-kit endpoint to net/http.
type Server struct {
	endpoint endpoint.Endpoint
	decode   DecodeRequestFunc
	encode   EncodeResponseFunc
}

// NewServer constructs an HTTP server for an endpoint.
func NewServer(endpoint endpoint.Endpoint, decode DecodeRequestFunc, encode EncodeResponseFunc) *Server {
	return &Server{endpoint: endpoint, decode: decode, encode: encode}
}

// ServeHTTP implements http.Handler.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	request, err := s.decode(r.Context(), r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := s.endpoint(r.Context(), request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := s.encode(r.Context(), w, response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
