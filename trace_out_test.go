package main

import (
	"os"
	"runtime/trace"
	"testing"
)

// 性能分析
func TestTraceOut(t *testing.T) {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	// TODO
}
