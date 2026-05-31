package service

import "context"

// Endpoint describes a legacy HTTP endpoint owned by this microservice.
type Endpoint struct {
	Name   string `json:"name"`
	Method string `json:"method"`
	Path   string `json:"path"`
}

// Service exposes notification-service business capabilities.
type Service interface {
	Health(ctx context.Context) (map[string]string, error)
	Endpoints(ctx context.Context) ([]Endpoint, error)
}

type service struct {
	name string
}

// New constructs the notification-service domain service.
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
	{Name: "API_SAVE_NOTIFICATIONGROUP", Method: "POST", Path: "/NotificationGroup/addnotificationgroup"},
	{Name: "API_UPDATE_NOTIFICATIONGROUP", Method: "PUT", Path: "/NotificationGroup/updatenotificationgroup"},
	{Name: "API_GET_NOTIFICATIONGROUPS", Method: "GET", Path: "/NotificationGroup/fetchnotificationgroupall"},
	{Name: "API_GET_NOTIFICATIONSBYCAMERAID", Method: "GET", Path: "/Notification/fetchnotificationbycameraid"},
	{Name: "API_GET_NOTIFICATIONBYID", Method: "GET", Path: "/Notification/fetchnotificationbyid"},
	{Name: "API_GET_NOTIFICATIONARRAY", Method: "GET", Path: "/Notification/getnotificationarray"},
	{Name: "API_GET_TODAYSNOTIFICATIONS", Method: "GET", Path: "/Notification/fetchnotification"},
	{Name: "API_POST_NOTIFICATION_STATUS", Method: "POST", Path: "/Notification/notificationstatusupdate"},
}
