package middleware

import (
	"strconv"
	"sync"
	"time"
)

// RateLimiter is a struct that holds the rate limiter.
type RateLimiter struct {
	Rate    time.Duration
	Burst   int
	mutex   sync.Mutex
	Limiter map[string]int
}

// NewRateLimit creates a new rate limiter.
func NewRateLimiter(rate time.Duration, burst string) *RateLimiter {
	convBurst, err := strconv.Atoi(burst)
	if err != nil {
		panic("Error converting burst to int")
	}

	return &RateLimiter{
		Rate:    rate,
		Burst:   convBurst,
		Limiter: make(map[string]int),
	}
}

// Limit is a middleware that limits the number of requests per second from a single IP.
func (r *RateLimiter) Limit(ip string) bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	count, ok := r.Limiter[ip]
	if !ok {
		r.Limiter[ip] = 1
		return true
	}
	if count >= r.Burst {
		return false
	}
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
