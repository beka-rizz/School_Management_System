package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func ResponseTimeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received Response from Response Time Middleware")
		start := time.Now()

		// Create a custom ResponseWriter to capture the status code
		wrappedWriter := &responseWriter{ResponseWriter: w, status: http.StatusOK}
		duration := time.Since(start)
		wrappedWriter.Header().Set("X-Response-Time", duration.String())
		next.ServeHTTP(wrappedWriter, r)
		// Calculate the duration
		duration = time.Since(start)
		// Log the request details
		fmt.Printf("Method: %s, URL: %s, Status: %d, Duration: %v\n", r.Method, r.URL, wrappedWriter.status, duration.String())
		fmt.Println("Sent Response from Response Time Middleware")
	})
}

// responseWriter
type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}
