package logger

import (
	"log/slog"
	"os"
)

// New returns a structured slog logger configured for service workloads.
func New(service string) *slog.Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true})
	return slog.New(handler).With("service", service)
}
