package main

import (
	"fmt"
	"net/http"
	"testing"
)

// File server
// Run the test server and send url "http://localhost:8080/" in your computer browser
func TestHttpFileServer(t *testing.T) {
	http.Handle("/", http.FileServer(http.Dir("/opt")))
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		fmt.Print(err)
	}
}
