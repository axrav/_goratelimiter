package main

import (
	"fmt"

	"github.com/axrav/rate_limit/internal/config"
	"github.com/axrav/rate_limit/internal/server"
)

func main() {
	// loading the config
	config.Load()
	fmt.Printf("[INFO] Starting the server on %s \n", config.Get("PORT"))
	// starting the server in a separate goroutine
	go server.Init()
	fmt.Println("[INFO] The server is ready to handle requests")
	// wait for the server to start

	select {}

}
