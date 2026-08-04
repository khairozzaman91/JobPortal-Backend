package middlewares

import (
	"net/http"
	"sync"
	"time"
)

type Visitor struct {
	Count     int
	LastVisit time.Time
}

type RateLimiter struct {
	mu       sync.Mutex
	visitors map[string]*Visitor
	limit    int
	window   time.Duration
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		visitors: make(map[string]*Visitor),
		limit:    limit,
		window:   window,
	}
}

func (rl *RateLimiter) Limit(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ip := r.RemoteAddr

		rl.mu.Lock()
		defer rl.mu.Unlock()

		visitor, exists := rl.visitors[ip]

		if !exists {
			rl.visitors[ip] = &Visitor{
				Count:     1,
				LastVisit: time.Now(),
			}

			next.ServeHTTP(w, r)
			return
		}

		if time.Since(visitor.LastVisit) > rl.window {
			visitor.Count = 1
			visitor.LastVisit = time.Now()

			next.ServeHTTP(w, r)
			return
		}

		if visitor.Count >= rl.limit {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		visitor.Count++
		visitor.LastVisit = time.Now()

		next.ServeHTTP(w, r)
	})
}