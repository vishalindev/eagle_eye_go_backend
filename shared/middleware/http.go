package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

// Chain composes HTTP middleware in declaration order.
func Chain(handler http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}
	return handler
}

// RequestLogger emits structured access logs for every request.
func RequestLogger(log *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			log.Info("request completed", "method", r.Method, "path", r.URL.Path, "duration", time.Since(start).String())
		})
	}
}
