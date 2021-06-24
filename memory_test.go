package main

import (
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"testing"
)

/**
terminal access command: go tool pprof http://localhost:8080/debug/pprof/heap
some commands:
top
*/

// access website: http://localhost:8080/debug/pprof
// grpc port is 9090 but web access need use 8080
func TestGrpcMemLeak(t *testing.T) {
	runtime.SetBlockProfileRate(1)
	go func() {
		http.ListenAndServe(":8080", nil)
	}()

	// new grpc server
	// newGrpcServer("127.0.0.1:9090")
}

// access website: http://localhost:9090/debug/pprof
func TestHttpMemLeak(t *testing.T) {
	runtime.SetBlockProfileRate(1)
	http.ListenAndServe(":9090", nil)

	// new grpc server
}
