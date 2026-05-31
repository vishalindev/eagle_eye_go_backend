package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"example.com/eagle-eye/services/admin-service/internal/service"
)

// Set contains go-kit endpoints for this service.
type Set struct {
	Health    endpoint.Endpoint
	Endpoints endpoint.Endpoint
}

// New wires domain service methods into go-kit endpoints.
func New(svc service.Service) Set {
	return Set{
		Health: func(ctx context.Context, request interface{}) (interface{}, error) {
			return svc.Health(ctx)
		},
		Endpoints: func(ctx context.Context, request interface{}) (interface{}, error) {
			return svc.Endpoints(ctx)
		},
	}
}
