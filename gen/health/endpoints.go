// Code generated by goa v3.1.3, DO NOT EDIT.
//
// health endpoints
//
// Command:
// $ goa gen github.com/wondenge/coop-notifiers/design

package health

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints wraps the "health" service endpoints.
type Endpoints struct {
	Show endpoint.Endpoint
}

// NewEndpoints wraps the methods of the "health" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Show: NewShowEndpoint(s),
	}
}

// Use applies the given middleware to all the "health" service endpoints.
func (e *Endpoints) Use(m func(endpoint.Endpoint) endpoint.Endpoint) {
	e.Show = m(e.Show)
}

// NewShowEndpoint returns an endpoint function that calls the method "show" of
// service "health".
func NewShowEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Show(ctx)
	}
}