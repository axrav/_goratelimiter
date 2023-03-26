package middleware

import (
	"sync"
	"time"
)

// RateLimiter is a struct that holds the rate limiter.
type RateLimiter struct {
	Rate    time.Duration
	Burst   int
	mutex   sync.RWMutex
	Limiter map[string]int
}

// NewRateLimit creates a new rate limiter.
func NewRateLimiter(rate time.Duration, burst int) *RateLimiter {
	return &RateLimiter{
		Rate:    rate,
		Burst:   burst,
		Limiter: make(map[string]int),
	}
}

// Limit is a middleware that limits the number of requests per second from a single IP.
func (r *RateLimiter) Limit(ip string) bool {
	r.mutex.RLock()
	count, ok := r.Limiter[ip]
	r.mutex.RUnlock()
	if !ok {
		r.mutex.Lock()
		defer r.mutex.Unlock()
		r.Limiter[ip] = 1
		return true
	}
	if count >= r.Burst {
		return false
	}
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.Limiter[ip]++
	return true
}

// Initiate is a goroutine that resets the rate limiter every second.
func (r *RateLimiter) Initiate() {
	timeTicker := time.NewTicker(r.Rate)
	defer timeTicker.Stop()

	for range timeTicker.C {
		r.mutex.Lock()
		for ip := range r.Limiter {
			r.Limiter[ip] = 0
		}
		r.mutex.Unlock()
	}
}
