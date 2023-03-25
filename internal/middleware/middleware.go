package middleware

import (
	"encoding/json"
	"net/http"
)

// rate limiting the requests to the server from a single IP using self written middleware

func RateLimitHandler(rl *RateLimiter, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !rl.Limit(r.RemoteAddr) {
			// if the limit is reached, return a 429 status code and response with a json message
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Too Many Requests",
			})

			return
		}
		next.ServeHTTP(w, r)
	})
}
