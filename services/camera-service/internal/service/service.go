package service

import "context"

// Endpoint describes a legacy HTTP endpoint owned by this microservice.
type Endpoint struct {
	Name   string `json:"name"`
	Method string `json:"method"`
	Path   string `json:"path"`
}

// Service exposes camera-service business capabilities.
type Service interface {
	Health(ctx context.Context) (map[string]string, error)
	Endpoints(ctx context.Context) ([]Endpoint, error)
}

type service struct {
	name string
}

// New constructs the camera-service domain service.
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
	{Name: "API_SAVE_CAMERA", Method: "POST", Path: "/Camera/addcamera"},
	{Name: "API_UPDATE_CAMERA", Method: "PUT", Path: "/Camera/updatecamera"},
	{Name: "API_GET_CAMERAS", Method: "GET", Path: "/Camera/fetchcamerasall"},
	{Name: "API_GET_CAMERA_UNITS", Method: "GET", Path: "/Camera/departmentwithcamera"},
}
