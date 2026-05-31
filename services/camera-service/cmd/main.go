package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"example.com/eagle-eye/services/camera-service/internal/endpoints"
	"example.com/eagle-eye/services/camera-service/internal/service"
	transporthttp "example.com/eagle-eye/services/camera-service/internal/transport/http"
	"example.com/eagle-eye/shared/logger"
	sharedmiddleware "example.com/eagle-eye/shared/middleware"
)

const serviceName = "camera-service"

func main() {
	log := logger.New(serviceName)
	svc := service.New(serviceName)
	endpointSet := endpoints.New(svc)
	handler := sharedmiddleware.Chain(transporthttp.NewHandler(endpointSet), sharedmiddleware.RequestLogger(log))

	server := &http.Server{
		Addr:              address(),
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		log.Info("starting service", "addr", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("service failed", slog.Any("error", err))
			os.Exit(1)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Error("graceful shutdown failed", slog.Any("error", err))
		os.Exit(1)
	}
	log.Info("service stopped")
}

func address() string {
	if value := os.Getenv("HTTP_ADDR"); value != "" {
		return value
	}
	return ":8080"
}
