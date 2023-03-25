package server

import (
	"flag"
	"net/http"
	"time"

	"github.com/axrav/rate_limit/internal/config"
	"github.com/axrav/rate_limit/internal/handlers"
	"github.com/axrav/rate_limit/internal/middleware"
)

// using http package

func Init() {
	// initializing the server here
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Online"))
	})
	// implementing the rate limiter
	n := flag.Int("N", 1, "Number of requests per second")
	flag.Parse()
	rateLimiter := middleware.NewRateLimiter(time.Second, *n)
	go rateLimiter.Initiate()
	// implementing the middleware and the handlers for the server
	http.Handle("/basic", middleware.RateLimitHandler(rateLimiter, http.HandlerFunc(handlers.BasicHandler)))
	// starting the server
	http.ListenAndServe(":"+config.Get("PORT"), nil)

}
