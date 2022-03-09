package main

import (
	_ "embed"
	"strings"
	"testing"
)

// go 1.16 speciality
//go:embed tmp/embed.txt
var str string

func TestEmbed(t *testing.T) {
	strArr := strings.Split(str, "\n")
	t.Log(len(strArr))
	for _, v := range strArr {
		t.Log(v)
	}
}
