package main

import (
	"net/http"
)

// HTTPResult - result of the http_request calls
type HTTPResult struct {
	URL        string
	StatusCode int
	Status     string
	Header     http.Header
	Content    []byte
}

// ConsulHealth - Consul health data
type ConsulHealth struct {
	Status []ConsulHealthStatus
}

// ConsulHealthStatus - Detailed health status
type ConsulHealthStatus struct {
	Node   string `json:"Node"`
	Status string `json:"Status"`
	Output string `json:"Output"`
}
