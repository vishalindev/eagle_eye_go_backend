package endpoint

import "context"

// Endpoint is the fundamental building block of go-kit services.
type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)
