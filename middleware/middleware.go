package middleware

import (
	"net/http"
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
	rateLimiter := NewRateLimiter(1000, time.Minute) // 1000 requests per minute
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rateLimiter.RateLimit(next.ServeHTTP).ServeHTTP(w, r)
	})
}
