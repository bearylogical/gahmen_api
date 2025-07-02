package middleware

import (
	"net/http"
	"os"
	"strconv"
	"time"
)

type Middleware func(http.Handler) http.Handler

func CreateStack(xs ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(xs) - 1; i >= 0; i-- {
			x := xs[i]
			next = x(next)
		}

		return next
	}
}

func RateLimit(next http.Handler) http.Handler {
	limit := 1000         // Default limit
	window := time.Minute // Default window

	if limitStr := os.Getenv("RATE_LIMIT_COUNT"); limitStr != "" {
		if parsedLimit, err := strconv.Atoi(limitStr); err == nil {
			limit = parsedLimit
		}
	}

	if windowStr := os.Getenv("RATE_LIMIT_WINDOW"); windowStr != "" {
		if parsedWindow, err := time.ParseDuration(windowStr); err == nil {
			window = parsedWindow
		}
	}

	rateLimiter := NewRateLimiter(limit, window)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rateLimiter.RateLimit(next.ServeHTTP).ServeHTTP(w, r)
	})
}
