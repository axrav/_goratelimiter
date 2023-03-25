package tests

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/axrav/rate_limit/internal/server"
)

func TestRateLimit(t *testing.T) {
	// setting the port
	os.Setenv("PORT", "8080")
	// starting the server
	go server.Init()
	// waiting for the server to start
	time.Sleep(1 * time.Second)
	// sending the request
	resp, err := http.Get("http://localhost:8080/basic")
	if err != nil {
		t.Error(err)
	}
	// checking the status code
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
	// checking the body
	var body map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	if body["message"] != "Hello Gopher!" {
		t.Errorf("Expected message 'Hello Gopher!', got '%s'", body["message"])
	}
	// sending the request again
	resp, err = http.Get("http://localhost:8080/basic")
	if err != nil {
		t.Error(err)
	}
	// checking the status code
	if resp.StatusCode != 429 {
		t.Errorf("Expected status code 429, got %d", resp.StatusCode)
	}
	// empty the body
	body = nil
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	if body["error"] != "Too Many Requests" {
		t.Errorf("Expected error message 'Too Many Requests', got '%s'", body["error"])
	}

}
