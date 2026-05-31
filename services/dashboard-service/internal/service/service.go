package service

import "context"

// Endpoint describes a legacy HTTP endpoint owned by this microservice.
type Endpoint struct {
	Name   string `json:"name"`
	Method string `json:"method"`
	Path   string `json:"path"`
}

// Service exposes dashboard-service business capabilities.
type Service interface {
	Health(ctx context.Context) (map[string]string, error)
	Endpoints(ctx context.Context) ([]Endpoint, error)
}

type service struct {
	name string
}

// New constructs the dashboard-service domain service.
func New(name string) Service {
	return service{name: name}
}

func (s service) Health(ctx context.Context) (map[string]string, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return map[string]string{"service": s.name, "status": "ok"}, nil
}

func (s service) Endpoints(ctx context.Context) ([]Endpoint, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return append([]Endpoint(nil), endpoints...), nil
}

var endpoints = []Endpoint{
	{Name: "API_GET_ACAMERAS", Method: "GET", Path: "/Dashboard/fetchanomalycameras"},
	{Name: "API_GET_UNITS", Method: "GET", Path: "/Dashboard/fetchunitsall"},
	{Name: "API_GET_USERTAB", Method: "GET", Path: "/Dashboard/fetchusertab"},
	{Name: "API_CREATE_USER_TAB", Method: "POST", Path: "/Dashboard/createusertab"},
	{Name: "API_FETCH_USER_TAB", Method: "GET", Path: "/Dashboard/fetchusertab"},
	{Name: "API_DELETE_TAB", Method: "DELETE", Path: "/Dashboard/deleteusertab/"},
}
