package clients

import "context"

// Dialer describes the internal gRPC client factory contract.
// Concrete generated clients can be registered here as protobuf stubs are added.
type Dialer interface {
	Close() error
}

// Factory centralizes internal service-client construction.
type Factory struct{}

// NewFactory returns a client factory for internal service-to-service calls.
func NewFactory() Factory { return Factory{} }

// Ping is a placeholder health check hook for future generated gRPC clients.
func (Factory) Ping(ctx context.Context) error { return ctx.Err() }
