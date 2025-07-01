package middleware

import (
	"net/http"
	"sync"
	"time"

	"gahmen-api/helpers"
)

type RateLimiter struct {
	mu      sync.Mutex
	clients map[string]*Client
	limit   int
	window  time.Duration
}

type Client struct {
	lastRequest time.Time
	requestCount int
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		clients: make(map[string]*Client),
		limit:   limit,
		window:  window,
	}
}

func (rl *RateLimiter) RateLimit(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr // Use remote address as a simple client identifier

		rl.mu.Lock()
		client, found := rl.clients[ip]
		if !found || time.Since(client.lastRequest) > rl.window {
			rl.clients[ip] = &Client{
				lastRequest: time.Now(),
				requestCount: 1,
			}
		} else {
			client.requestCount++
				rl.clients[ip] = client
		}

		if rl.clients[ip].requestCount > rl.limit {
			rl.mu.Unlock()
			helpers.WriteJSON(w, http.StatusTooManyRequests, map[string]string{"error": "Too many requests"})
			return
		}
		rl.mu.Unlock()

		next(w, r)
	}
}
