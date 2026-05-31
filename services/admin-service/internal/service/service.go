package service

import "context"

// Endpoint describes a legacy HTTP endpoint owned by this microservice.
type Endpoint struct {
	Name   string `json:"name"`
	Method string `json:"method"`
	Path   string `json:"path"`
}

// Service exposes admin-service business capabilities.
type Service interface {
	Health(ctx context.Context) (map[string]string, error)
	Endpoints(ctx context.Context) ([]Endpoint, error)
}

type service struct {
	name string
}

// New constructs the admin-service domain service.
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
	{Name: "API_SAVE_EMPLOYEE", Method: "POST", Path: "/Employee/createuser"},
	{Name: "API_UPDATE_EMPLOYEE", Method: "PUT", Path: "/Employee/updateuser"},
	{Name: "API_GET_EMPLOYEES", Method: "GET", Path: "/Employee/fetchusers"},
	{Name: "API_SAVE_PLANTSETUP", Method: "POST", Path: "/Plant/createplant"},
	{Name: "API_UPDATE_PLANTSETUP", Method: "PUT", Path: "/Plant/updateplant"},
	{Name: "API_GET_PLANTSETUPS", Method: "GET", Path: "/Plant/fetchplant"},
}
