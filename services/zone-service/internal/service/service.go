package service

import "context"

// Endpoint describes a legacy HTTP endpoint owned by this microservice.
type Endpoint struct {
	Name   string `json:"name"`
	Method string `json:"method"`
	Path   string `json:"path"`
}

// Service exposes zone-service business capabilities.
type Service interface {
	Health(ctx context.Context) (map[string]string, error)
	Endpoints(ctx context.Context) ([]Endpoint, error)
}

type service struct {
	name string
}

// New constructs the zone-service domain service.
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
	{Name: "API_SAVE_ZONE", Method: "POST", Path: "/Zone/createzone"},
	{Name: "API_UPDATE_ZONE", Method: "PUT", Path: "/Zone/updatezone"},
	{Name: "API_GET_ZONES", Method: "GET", Path: "/Zone/fetchzoneall"},
	{Name: "API_GET_CAMERABYZONEID", Method: "GET", Path: "/Zone/fetchzonebyid"},
	{Name: "API_GET_ZONE_CODE", Method: "GET", Path: "/Zone/createzonecode"},
	{Name: "API_UPDATE_ZONE_CAMERA", Method: "PUT", Path: "/Zone/zonecameramap"},
	{Name: "API_UPDATE_ZONE_RULE", Method: "PUT", Path: "/Zone/zonerulemap"},
	{Name: "API_GET_ZONE_USER_TAB", Method: "GET", Path: "/Zone/fetchzonebytabid"},
}
