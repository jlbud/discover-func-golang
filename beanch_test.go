package main

import (
	"bytes"
	"strings"
	"testing"
)

var (
	ss = "1234567890abcdefghijklmnopqrstuvwxyz"
	bs = []byte(ss)
	rn = 'a'
	bt = byte('a')
)

func BenchmarkBuilderWrite(b *testing.B) {
	var builder strings.Builder
	for i := 0; i < b.N; i++ {
		builder.Write(bs)
	}
}

func BenchmarkBuilderWriteByte(b *testing.B) {
	var builder strings.Builder
	for i := 0; i < b.N; i++ {
		builder.WriteByte(bt)
	}
}

func BenchmarkBuilderWriteRune(b *testing.B) {
	var builder strings.Builder
	for i := 0; i < b.N; i++ {
		builder.WriteRune(rn)
	}
}

func BenchmarkBuilderWriteString(b *testing.B) {
	var builder strings.Builder
	for i := 0; i < b.N; i++ {
		builder.WriteString(ss)
	}
}

func BenchmarkBufferWrite(b *testing.B) {
	var buffer bytes.Buffer
	for i := 0; i < b.N; i++ {
		buffer.Write(bs)
	}
}

func BenchmarkBufferWriteByte(b *testing.B) {
	var buffer bytes.Buffer
	for i := 0; i < b.N; i++ {
		buffer.WriteByte(bt)
	}
}

func BenchmarkBufferWriteRune(b *testing.B) {
	var buffer bytes.Buffer
	for i := 0; i < b.N; i++ {
		buffer.WriteRune(rn)
	}
}

func BenchmarkBufferWriteString(b *testing.B) {
	var buffer bytes.Buffer
	for i := 0; i < b.N; i++ {
		buffer.WriteString(ss)
	}
}
